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
	MonitorDuration time.Duration `json:"monitor_duration"`
	RequestInterval time.Duration `json:"request_interval"`
	ConcurrentReqs  int           `json:"concurrent_requests"`
	Timeout         time.Duration `json:"timeout"`
	Method          string        `json:"method"`
	Headers         []Header      `json:"headers"`
	Body            string        `json:"body"`
	OutputFormat    string        `json:"output_format"`
	LogLevel        string        `json:"log_level"`
	ReportInterval  time.Duration `json:"report_interval"`
	SaveToFile      bool          `json:"save_to_file"`
	OutputFile      string        `json:"output_file"`
}

// Header represents an HTTP header
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// RequestResult holds individual request results
type RequestResult struct {
	Timestamp    time.Time     `json:"timestamp"`
	ResponseTime time.Duration `json:"response_time"`
	StatusCode   int           `json:"status_code"`
	Success      bool          `json:"success"`
	Error        string        `json:"error,omitempty"`
}

// TimeWindowStats holds statistics for a specific time window
type TimeWindowStats struct {
	WindowStart      time.Time     `json:"window_start"`
	WindowEnd        time.Time     `json:"window_end"`
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
	StatusCodes      map[int]int   `json:"status_codes"`
	SuccessRate      float64       `json:"success_rate"`
}

// ContinuousAPIMonitor handles continuous monitoring operations
type ContinuousAPIMonitor struct {
	config       *Config
	client       *http.Client
	results      []RequestResult
	mutex        sync.RWMutex
	logger       *log.Logger
	startTime    time.Time
	stopChan     chan struct{}
}

// NewContinuousAPIMonitor creates a new continuous API monitor instance
func NewContinuousAPIMonitor(config *Config) *ContinuousAPIMonitor {
	return &ContinuousAPIMonitor{
		config: config,
		client: &http.Client{
			Timeout: config.Timeout,
			Transport: &http.Transport{
				MaxIdleConns:       100,
				IdleConnTimeout:    90 * time.Second,
				DisableCompression: false,
			},
		},
		results:   make([]RequestResult, 0),
		logger:    log.New(os.Stdout, "[CONTINUOUS-API-MONITOR] ", log.LstdFlags|log.Lshortfile),
		startTime: time.Now(),
		stopChan:  make(chan struct{}),
	}
}

// makeRequest performs a single HTTP request and records the result
func (cam *ContinuousAPIMonitor) makeRequest(ctx context.Context) {
	req, err := http.NewRequestWithContext(ctx, cam.config.Method, cam.config.URL, nil)
	if err != nil {
		cam.recordResult(RequestResult{
			Timestamp: time.Now(),
			Success:   false,
			Error:     fmt.Sprintf("creating request: %v", err),
		})
		return
	}

	// Add custom headers
	for _, header := range cam.config.Headers {
		req.Header.Set(header.Key, header.Value)
	}

	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "Continuous-API-Monitor/1.0")
	}

	start := time.Now()
	resp, err := cam.client.Do(req)
	responseTime := time.Since(start)
	timestamp := time.Now()

	result := RequestResult{
		Timestamp:    timestamp,
		ResponseTime: responseTime,
	}

	if err != nil {
		result.Success = false
		result.Error = err.Error()
		cam.logger.Printf("Request failed: %v (Response time: %v)", err, responseTime)
	} else {
		defer resp.Body.Close()
		result.Success = true
		result.StatusCode = resp.StatusCode
		
		if cam.config.LogLevel == "debug" {
			cam.logger.Printf("Request successful: Status %d, Response time: %v", resp.StatusCode, responseTime)
		}
	}

	cam.recordResult(result)
}

// recordResult safely adds a result to the results slice
func (cam *ContinuousAPIMonitor) recordResult(result RequestResult) {
	cam.mutex.Lock()
	defer cam.mutex.Unlock()
	cam.results = append(cam.results, result)
}

// cleanOldResults removes results older than the monitor duration
func (cam *ContinuousAPIMonitor) cleanOldResults() {
	cam.mutex.Lock()
	defer cam.mutex.Unlock()

	cutoff := time.Now().Add(-cam.config.MonitorDuration)
	
	// Find the first index that's still within the window
	firstValidIndex := 0
	for i, result := range cam.results {
		if result.Timestamp.After(cutoff) {
			firstValidIndex = i
			break
		}
	}

	// Remove old results
	if firstValidIndex > 0 {
		cam.results = cam.results[firstValidIndex:]
		if cam.config.LogLevel == "debug" {
			cam.logger.Printf("Cleaned %d old results", firstValidIndex)
		}
	}
}

// getWindowStats calculates statistics for the specified time window
func (cam *ContinuousAPIMonitor) getWindowStats(windowDuration time.Duration) *TimeWindowStats {
	cam.mutex.RLock()
	defer cam.mutex.RUnlock()

	now := time.Now()
	windowStart := now.Add(-windowDuration)
	
	// Filter results within the window
	var windowResults []RequestResult
	for _, result := range cam.results {
		if result.Timestamp.After(windowStart) {
			windowResults = append(windowResults, result)
		}
	}

	if len(windowResults) == 0 {
		return &TimeWindowStats{
			WindowStart:   windowStart,
			WindowEnd:     now,
			URL:           cam.config.URL,
			StatusCodes:   make(map[int]int),
		}
	}

	stats := &TimeWindowStats{
		WindowStart: windowStart,
		WindowEnd:   now,
		URL:         cam.config.URL,
		StatusCodes: make(map[int]int),
	}

	var responseTimes []time.Duration
	var totalResponseTime time.Duration

	for _, result := range windowResults {
		stats.TotalRequests++
		
		if result.Success {
			stats.SuccessfulReqs++
			stats.StatusCodes[result.StatusCode]++
			responseTimes = append(responseTimes, result.ResponseTime)
			totalResponseTime += result.ResponseTime
		} else {
			stats.FailedReqs++
		}
	}

	if len(responseTimes) > 0 {
		// Sort for percentile calculations
		sort.Slice(responseTimes, func(i, j int) bool {
			return responseTimes[i] < responseTimes[j]
		})

		stats.MinResponseTime = responseTimes[0]
		stats.MaxResponseTime = responseTimes[len(responseTimes)-1]
		stats.AvgResponseTime = totalResponseTime / time.Duration(len(responseTimes))
		stats.MedianTime = calculatePercentile(responseTimes, 50)
		stats.P95Time = calculatePercentile(responseTimes, 95)
		stats.P99Time = calculatePercentile(responseTimes, 99)
	}

	if stats.TotalRequests > 0 {
		stats.SuccessRate = float64(stats.SuccessfulReqs) / float64(stats.TotalRequests) * 100
	}

	return stats
}

// calculatePercentile calculates the specified percentile from sorted response times
func calculatePercentile(sortedTimes []time.Duration, percentile int) time.Duration {
	if len(sortedTimes) == 0 {
		return 0
	}
	
	index := int(float64(len(sortedTimes)) * float64(percentile) / 100.0)
	if index >= len(sortedTimes) {
		index = len(sortedTimes) - 1
	}
	return sortedTimes[index]
}

// startMonitoring begins the continuous monitoring process
func (cam *ContinuousAPIMonitor) startMonitoring(ctx context.Context) {
	cam.logger.Printf("Starting continuous monitoring for %v on %s", cam.config.MonitorDuration, cam.config.URL)
	cam.logger.Printf("Request interval: %v, Concurrent requests: %d", cam.config.RequestInterval, cam.config.ConcurrentReqs)

	// Semaphore for controlling concurrency
	semaphore := make(chan struct{}, cam.config.ConcurrentReqs)
	
	// Ticker for making requests
	requestTicker := time.NewTicker(cam.config.RequestInterval)
	defer requestTicker.Stop()

	// Ticker for periodic reports
	var reportTicker *time.Ticker
	if cam.config.ReportInterval > 0 {
		reportTicker = time.NewTicker(cam.config.ReportInterval)
		defer reportTicker.Stop()
	}

	// Ticker for cleanup
	cleanupTicker := time.NewTicker(1 * time.Minute)
	defer cleanupTicker.Stop()

	// Monitor until duration expires or context is cancelled
	endTime := time.Now().Add(cam.config.MonitorDuration)

	for {
		select {
		case <-ctx.Done():
			cam.logger.Println("Monitoring stopped by context cancellation")
			return
		case <-cam.stopChan:
			cam.logger.Println("Monitoring stopped by stop signal")
			return
		case <-requestTicker.C:
			if time.Now().After(endTime) {
				cam.logger.Printf("Monitoring duration (%v) completed", cam.config.MonitorDuration)
				return
			}
			
			select {
			case semaphore <- struct{}{}:
				go func() {
					defer func() { <-semaphore }()
					cam.makeRequest(ctx)
				}()
			default:
				// All workers busy, skip this tick
			}
		case <-cleanupTicker.C:
			cam.cleanOldResults()
		case <-reportTicker.C:
			if reportTicker != nil {
				cam.printPeriodicReport()
			}
		}
	}
}

// printPeriodicReport prints a periodic status report
func (cam *ContinuousAPIMonitor) printPeriodicReport() {
	stats := cam.getWindowStats(cam.config.MonitorDuration)
	
	cam.logger.Printf("Periodic Report - Total requests: %d, Success rate: %.2f%%, Avg response time: %v",
		stats.TotalRequests, stats.SuccessRate, stats.AvgResponseTime)
}

// printFinalResults prints the final monitoring results
func (cam *ContinuousAPIMonitor) printFinalResults() {
	stats := cam.getWindowStats(cam.config.MonitorDuration)

	switch cam.config.OutputFormat {
	case "json":
		cam.printJSONResults(stats)
	default:
		cam.printTableResults(stats)
	}

	// Save to file if configured
	if cam.config.SaveToFile && cam.config.OutputFile != "" {
		cam.saveResultsToFile(stats)
	}
}

// printTableResults prints results in a formatted table
func (cam *ContinuousAPIMonitor) printTableResults(stats *TimeWindowStats) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Continuous API Response Time Monitoring Results")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("URL: %s\n", stats.URL)
	fmt.Printf("Monitoring Window: %v to %v\n", stats.WindowStart.Format("2006-01-02 15:04:05"), stats.WindowEnd.Format("2006-01-02 15:04:05"))
	fmt.Printf("Window Duration: %v\n", stats.WindowEnd.Sub(stats.WindowStart))
	fmt.Printf("Total Requests: %d\n", stats.TotalRequests)
	fmt.Printf("Successful: %d (%.2f%%)\n", stats.SuccessfulReqs, stats.SuccessRate)
	fmt.Printf("Failed: %d (%.2f%%)\n", stats.FailedReqs, 100-stats.SuccessRate)
	
	if stats.SuccessfulReqs > 0 {
		fmt.Println("\nResponse Time Statistics:")
		fmt.Printf("  Minimum:     %v\n", stats.MinResponseTime)
		fmt.Printf("  Maximum:     %v\n", stats.MaxResponseTime)
		fmt.Printf("  Average:     %v\n", stats.AvgResponseTime)
		fmt.Printf("  Median:      %v\n", stats.MedianTime)
		fmt.Printf("  95th %%ile:   %v\n", stats.P95Time)
		fmt.Printf("  99th %%ile:   %v\n", stats.P99Time)

		fmt.Println("\nHTTP Status Codes:")
		for code, count := range stats.StatusCodes {
			percentage := float64(count) / float64(stats.SuccessfulReqs) * 100
			fmt.Printf("  %d: %d (%.2f%%)\n", code, count, percentage)
		}
	}
	fmt.Println(strings.Repeat("=", 80))
}

// printJSONResults prints results in JSON format
func (cam *ContinuousAPIMonitor) printJSONResults(stats *TimeWindowStats) {
	jsonData, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		cam.logger.Printf("Error marshaling JSON: %v", err)
		return
	}
	fmt.Println(string(jsonData))
}

// saveResultsToFile saves results to a file
func (cam *ContinuousAPIMonitor) saveResultsToFile(stats *TimeWindowStats) {
	file, err := os.Create(cam.config.OutputFile)
	if err != nil {
		cam.logger.Printf("Error creating output file: %v", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	
	if err := encoder.Encode(stats); err != nil {
		cam.logger.Printf("Error writing to output file: %v", err)
		return
	}

	cam.logger.Printf("Results saved to %s", cam.config.OutputFile)
}

// Stop gracefully stops the monitoring
func (cam *ContinuousAPIMonitor) Stop() {
	close(cam.stopChan)
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

// getDefaultConfig returns a default configuration for 72-hour monitoring
func getDefaultConfig() *Config {
	return &Config{
		URL:             "https://httpbin.org/delay/1",
		MonitorDuration: 72 * time.Hour,
		RequestInterval: 30 * time.Second,
		ConcurrentReqs:  1,
		Timeout:         30 * time.Second,
		Method:          "GET",
		Headers:         []Header{},
		OutputFormat:    "table",
		LogLevel:        "info",
		ReportInterval:  1 * time.Hour,
		SaveToFile:      false,
		OutputFile:      "api_monitor_results.json",
	}
}

func main() {
	var (
		url             = flag.String("url", "", "API endpoint URL to monitor")
		duration        = flag.Duration("duration", 72*time.Hour, "Total monitoring duration")
		interval        = flag.Duration("interval", 30*time.Second, "Interval between requests")
		concurrent      = flag.Int("concurrent", 1, "Number of concurrent requests")
		timeout         = flag.Duration("timeout", 30*time.Second, "Request timeout")
		method          = flag.String("method", "GET", "HTTP method")
		outputFormat    = flag.String("output", "table", "Output format (table/json)")
		configFile      = flag.String("config", "", "Configuration file path")
		logLevel        = flag.String("log-level", "info", "Log level (info/debug)")
		reportInterval  = flag.Duration("report-interval", 1*time.Hour, "Interval for periodic reports (0 to disable)")
		saveToFile      = flag.Bool("save", false, "Save results to file")
		outputFile      = flag.String("output-file", "api_monitor_results.json", "Output file path")
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
		config.MonitorDuration = *duration
		config.RequestInterval = *interval
		config.ConcurrentReqs = *concurrent
		config.Timeout = *timeout
		config.Method = *method
		config.OutputFormat = *outputFormat
		config.LogLevel = *logLevel
		config.ReportInterval = *reportInterval
		config.SaveToFile = *saveToFile
		config.OutputFile = *outputFile
	}

	// Validate required parameters
	if config.URL == "" {
		log.Fatal("URL is required. Use -url flag or provide config file.")
	}

	if config.MonitorDuration <= 0 {
		log.Fatal("Monitor duration must be greater than 0")
	}

	if config.ConcurrentReqs <= 0 {
		log.Fatal("Concurrent requests must be greater than 0")
	}

	// Create monitor
	monitor := NewContinuousAPIMonitor(config)

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		monitor.logger.Println("Received shutdown signal, stopping...")
		monitor.Stop()
		cancel()
	}()

	// Run monitoring
	monitor.startMonitoring(ctx)

	// Print final results
	monitor.printFinalResults()
}
