# If Prometheus CRDs are installed, proceed:
# Create ServiceMonitor for Prometheus
cat <<EOF | kubectl apply -f -
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: postgresql-metrics
  namespace: postgresql
  labels:
    release: prometheus
spec:
  selector:
    matchLabels:
      app: postgresql
  endpoints:
    - port: metrics
      interval: 30s
      scrapeTimeout: 10s
EOF

# Create PrometheusRule for alerts
cat <<EOF | kubectl apply -f -
apiVersion: monitoring.coreos.com/v1
kind: PrometheusRule
metadata:
  name: postgresql-alerts
  namespace: postgresql
  labels:
    prometheus: k8s
    role: alert-rules
spec:
  groups:
    - name: postgresql
      rules:
        - alert: PostgreSQLHighConnectionCount
          expr: sum by (pod) (pg_stat_activity_count{datname!~"template.*|postgres"}) > 100
          for: 5m
          labels:
            severity: warning
          annotations:
            summary: "PostgreSQL high connection count"
            description: "PostgreSQL instance has {{ \$value }} connections"
        - alert: PostgreSQLDown
          expr: pg_up == 0
          for: 1m
          labels:
            severity: critical
          annotations:
            summary: "PostgreSQL down"
            description: "PostgreSQL instance is down"
        - alert: PostgreSQLHighDiskUsage
          expr: (pg_database_size_bytes / 1024 / 1024 / 1024) > 40
          for: 5m
          labels:
            severity: warning
          annotations:
            summary: "PostgreSQL high disk usage"
            description: "PostgreSQL database size exceeds 40GB"
EOF
