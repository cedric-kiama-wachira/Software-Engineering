kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: postgres-init-runner
  namespace: postgresql
data:
  init-db.sh: |
    #!/bin/bash
    set -e
    
    # Wait for PostgreSQL to be ready
    until pg_isready -U postgres; do
      echo "Waiting for PostgreSQL to be ready..."
      sleep 2
    done
    
    echo "PostgreSQL is ready, initializing extensions..."
    psql -U postgres -f /docker-entrypoint-initdb.d/init-extensions.sql
    echo "Extensions initialization completed."
EOF
