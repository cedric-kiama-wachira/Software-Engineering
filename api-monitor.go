package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"
)

// Config holds the configuration for the API monitor
type Config struct {
	URL             string        `json:"url"`
	RequestCount    int           `json:"request_count"`
	ConcurrentReqs  int           `json:"concurrent_requests"`
	Timeout         time.Duration `json:"timeout"`
	Interval        time.Duration `json:"interval"`
	Method          string        `json:"method"`
	Headers         []Header      `json:"headers"`
	Body            string        `json:"body"`
	OutputFormat    string        `json:"output_format"`
	LogLevel        string        `json:"log_level"`
}

// Header represents an HTTP header
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ResponseStats holds the response time statistics
type ResponseStats struct {
	URL              string        `json:"url"`
	TotalRequests    int           `json:"total_requests"`
	SuccessfulReqs   int           `json:"successful_requests"`
	FailedReqs       int           `json:"failed_requests"`
	MinResponseTime  time.Duration `json:"min_response_time"`
	MaxResponseTime  time.Duration `json:"max_response_time"`
	AvgResponseTime  time.Duration `json:"avg_response_time"`
	MedianTime       time.Duration `json:"median_response_time"`
	P95Time          time.Duration `json:"p95_response_time"`
	P99Time          time.Duration `json:"p99_response_time"`
	ResponseTimes    []time.Duration `json:"-"`
	StatusCodes      map[int]int   `json:"status_codes"`
	StartTime        time.Time     `json:"start_time"`
	EndTime          time.Time     `json:"end_time"`
	TotalDuration    time.Duration `json:"total_duration"`
}

// APIMonitor handles the monitoring operations
type APIMonitor struct {
	config *Config
	client *http.Client
	stats  *ResponseStats
	mutex  sync.RWMutex
	logger *log.Logger
}

// NewAPIMonitor creates a new API monitor instance
func NewAPIMonitor(config *Config) *APIMonitor {
	return &APIMonitor{
		config: config,
		client: &http.Client{
			Timeout: config.Timeout,
			Transport: &http.Transport{
				MaxIdleConns:       100,
				IdleConnTimeout:    90 * time.Second,
				DisableCompression: false,
			},
		},
		stats: &ResponseStats{
			URL:           config.URL,
			StatusCodes:   make(map[int]int),
			ResponseTimes: make([]time.Duration, 0),
		},
		logger: log.New(os.Stdout, "[API-MONITOR] ", log.LstdFlags|log.Lshortfile),
	}
}

// makeRequest performs a single HTTP request and measures response time
func (am *APIMonitor) makeRequest(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, am.config.Method, am.config.URL, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	// Add custom headers
	for _, header := range am.config.Headers {
		req.Header.Set(header.Key, header.Value)
	}

	// Set default User-Agent if not provided
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "API-Monitor/1.0")
	}

	start := time.Now()
	resp, err := am.client.Do(req)
	responseTime := time.Since(start)

	am.mutex.Lock()
	defer am.mutex.Unlock()

	am.stats.TotalRequests++
	am.stats.ResponseTimes = append(am.stats.ResponseTimes, responseTime)

	if err != nil {
		am.stats.FailedReqs++
		am.logger.Printf("Request failed: %v (Response time: %v)", err, responseTime)
		return err
	}
	defer resp.Body.Close()

	am.stats.SuccessfulReqs++
	am.stats.StatusCodes[resp.StatusCode]++

	// Log successful requests in debug mode
	if am.config.LogLevel == "debug" {
		am.logger.Printf("Request successful: Status %d, Response time: %v", resp.StatusCode, responseTime)
	}

	return nil
}

// runConcurrentRequests executes requests concurrently
func (am *APIMonitor) runConcurrentRequests(ctx context.Context) {
	am.stats.StartTime = time.Now()
	defer func() {
		am.stats.EndTime = time.Now()
		am.stats.TotalDuration = am.stats.EndTime.Sub(am.stats.StartTime)
	}()

	semaphore := make(chan struct{}, am.config.ConcurrentReqs)
	var wg sync.WaitGroup

	requestsPerformed := 0
	ticker := time.NewTicker(am.config.Interval)
	defer ticker.Stop()

	am.logger.Printf("Starting API monitoring: %d requests to %s", am.config.RequestCount, am.config.URL)

	for requestsPerformed < am.config.RequestCount {
		select {
		case <-ctx.Done():
			am.logger.Println("Monitoring stopped by context cancellation")
			wg.Wait()
			return
		case <-ticker.C:
			select {
			case semaphore <- struct{}{}:
				wg.Add(1)
				go func() {
					defer wg.Done()
					defer func() { <-semaphore }()
					
					if err := am.makeRequest(ctx); err != nil {
						if am.config.LogLevel == "debug" {
							am.logger.Printf("Request error: %v", err)
						}
					}
				}()
				requestsPerformed++
			default:
				// All workers busy, continue to next tick
			}
		}
	}

	wg.Wait()
	am.logger.Println("All requests completed")
}

// calculateStats computes the response time statistics
func (am *APIMonitor) calculateStats() {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	if len(am.stats.ResponseTimes) == 0 {
		return
	}

	// Sort response times for percentile calculations
	sort.Slice(am.stats.ResponseTimes, func(i, j int) bool {
		return am.stats.ResponseTimes[i] < am.stats.ResponseTimes[j]
	})

	am.stats.MinResponseTime = am.stats.ResponseTimes[0]
	am.stats.MaxResponseTime = am.stats.ResponseTimes[len(am.stats.ResponseTimes)-1]

	// Calculate average
	var total time.Duration
	for _, rt := range am.stats.ResponseTimes {
		total += rt
	}
	am.stats.AvgResponseTime = total / time.Duration(len(am.stats.ResponseTimes))

	// Calculate percentiles
	am.stats.MedianTime = am.calculatePercentile(50)
	am.stats.P95Time = am.calculatePercentile(95)
	am.stats.P99Time = am.calculatePercentile(99)
}

// calculatePercentile calculates the specified percentile
func (am *APIMonitor) calculatePercentile(percentile int) time.Duration {
	if len(am.stats.ResponseTimes) == 0 {
		return 0
	}
	
	index := int(float64(len(am.stats.ResponseTimes)) * float64(percentile) / 100.0)
	if index >= len(am.stats.ResponseTimes) {
		index = len(am.stats.ResponseTimes) - 1
	}
	return am.stats.ResponseTimes[index]
}

// printResults outputs the monitoring results
func (am *APIMonitor) printResults() {
	am.calculateStats()

	switch am.config.OutputFormat {
	case "json":
		am.printJSONResults()
	default:
		am.printTableResults()
	}
}

// printTableResults prints results in a formatted table
func (am *APIMonitor) printTableResults() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("API Response Time Monitoring Results")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("URL: %s\n", am.stats.URL)
	fmt.Printf("Total Duration: %v\n", am.stats.TotalDuration)
	fmt.Printf("Total Requests: %d\n", am.stats.TotalRequests)
	fmt.Printf("Successful: %d (%.2f%%)\n", am.stats.SuccessfulReqs, 
		float64(am.stats.SuccessfulReqs)/float64(am.stats.TotalRequests)*100)
	fmt.Printf("Failed: %d (%.2f%%)\n", am.stats.FailedReqs,
		float64(am.stats.FailedReqs)/float64(am.stats.TotalRequests)*100)
	
	fmt.Println("\nResponse Time Statistics:")
	fmt.Printf("  Minimum:     %v\n", am.stats.MinResponseTime)
	fmt.Printf("  Maximum:     %v\n", am.stats.MaxResponseTime)
	fmt.Printf("  Average:     %v\n", am.stats.AvgResponseTime)
	fmt.Printf("  Median:      %v\n", am.stats.MedianTime)
	fmt.Printf("  95th %%ile:   %v\n", am.stats.P95Time)
	fmt.Printf("  99th %%ile:   %v\n", am.stats.P99Time)

	fmt.Println("\nHTTP Status Codes:")
	for code, count := range am.stats.StatusCodes {
		percentage := float64(count) / float64(am.stats.SuccessfulReqs) * 100
		fmt.Printf("  %d: %d (%.2f%%)\n", code, count, percentage)
	}
	fmt.Println(strings.Repeat("=", 80))
}

// printJSONResults prints results in JSON format
func (am *APIMonitor) printJSONResults() {
	jsonData, err := json.MarshalIndent(am.stats, "", "  ")
	if err != nil {
		am.logger.Printf("Error marshaling JSON: %v", err)
		return
	}
	fmt.Println(string(jsonData))
}

// loadConfig loads configuration from a file
func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// getDefaultConfig returns a default configuration
func getDefaultConfig() *Config {
	return &Config{
		URL:            "https://httpbin.org/delay/1",
		RequestCount:   10,
		ConcurrentReqs: 1,
		Timeout:        30 * time.Second,
		Interval:       100 * time.Millisecond,
		Method:         "GET",
		Headers:        []Header{},
		OutputFormat:   "table",
		LogLevel:       "info",
	}
}

func main() {
	var (
		url            = flag.String("url", "", "API endpoint URL to monitor")
		count          = flag.Int("count", 10, "Number of requests to make")
		concurrent     = flag.Int("concurrent", 1, "Number of concurrent requests")
		timeout        = flag.Duration("timeout", 30*time.Second, "Request timeout")
		interval       = flag.Duration("interval", 100*time.Millisecond, "Interval between requests")
		method         = flag.String("method", "GET", "HTTP method")
		outputFormat   = flag.String("output", "table", "Output format (table/json)")
		configFile     = flag.String("config", "", "Configuration file path")
		logLevel       = flag.String("log-level", "info", "Log level (info/debug)")
	)
	flag.Parse()

	var config *Config
	var err error

	// Load config from file if provided, otherwise use flags/defaults
	if *configFile != "" {
		config, err = loadConfig(*configFile)
		if err != nil {
			log.Fatalf("Error loading config file: %v", err)
		}
	} else {
		config = getDefaultConfig()
		
		// Override with command line flags if provided
		if *url != "" {
			config.URL = *url
		}
		config.RequestCount = *count
		config.ConcurrentReqs = *concurrent
		config.Timeout = *timeout
		config.Interval = *interval
		config.Method = *method
		config.OutputFormat = *outputFormat
		config.LogLevel = *logLevel
	}

	// Validate required parameters
	if config.URL == "" {
		log.Fatal("URL is required. Use -url flag or provide config file.")
	}

	if config.RequestCount <= 0 {
		log.Fatal("Request count must be greater than 0")
	}

	if config.ConcurrentReqs <= 0 {
		log.Fatal("Concurrent requests must be greater than 0")
	}

	// Create monitor
	monitor := NewAPIMonitor(config)

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		monitor.logger.Println("Received shutdown signal, stopping...")
		cancel()
	}()

	// Run monitoring
	monitor.runConcurrentRequests(ctx)

	// Print results
	monitor.printResults()
}
