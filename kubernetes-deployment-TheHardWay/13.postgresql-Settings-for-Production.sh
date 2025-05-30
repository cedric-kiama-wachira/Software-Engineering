cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: postgresql-config
  namespace: postgresql
data:
  postgresql.conf: |
    # Memory settings optimized for your 16GB worker nodes
    shared_buffers = '4GB'
    effective_cache_size = '8GB'
    work_mem = '32MB'
    maintenance_work_mem = '512MB'
    
    # Write-ahead log settings
    wal_level = logical
    max_wal_size = '2GB'
    min_wal_size = '1GB'
    checkpoint_completion_target = 0.9
    
    # Connection settings
    max_connections = 200
    
    # Query tuning
    random_page_cost = 1.1
    effective_io_concurrency = 200
    
    # TimescaleDB settings
    timescaledb.max_background_workers = 8
    
    # Logging settings
    log_min_duration_statement = 1000
    log_checkpoints = on
    log_connections = on
    log_disconnections = on
    log_lock_waits = on
    log_temp_files = 0
EOF
