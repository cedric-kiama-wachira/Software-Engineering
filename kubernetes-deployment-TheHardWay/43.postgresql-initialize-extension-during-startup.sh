kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: postgres-init-extensions
  namespace: postgresql
data:
  init-extensions.sh: |
    #!/bin/bash
    set -e
    
    echo "Initializing PostgreSQL extensions..."
    
    # Wait for PostgreSQL to become ready
    until pg_isready -U postgres; do
      echo "Waiting for PostgreSQL to be ready..."
      sleep 5
    done
    
    # Install all extensions
    psql -U postgres -c "
      CREATE EXTENSION IF NOT EXISTS timescaledb;
      CREATE EXTENSION IF NOT EXISTS pgpcre;
      CREATE EXTENSION IF NOT EXISTS postgis;
      CREATE EXTENSION IF NOT EXISTS hstore;
      CREATE EXTENSION IF NOT EXISTS pg_stat_statements;
      CREATE EXTENSION IF NOT EXISTS pgcrypto;
      CREATE EXTENSION IF NOT EXISTS pg_partman;
      CREATE EXTENSION IF NOT EXISTS postgres_fdw;
      CREATE EXTENSION IF NOT EXISTS vector;
      CREATE EXTENSION IF NOT EXISTS timescaledb_toolkit;
    "
    
    echo "Extensions initialization completed successfully."
EOF
