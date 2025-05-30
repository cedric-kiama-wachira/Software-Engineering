# Generate a secure password
PASSWORD=$(openssl rand -base64 20)
echo "Your PostgreSQL password is: $PASSWORD"
echo "Make sure to save this in a secure place!"

# Create Secret for PostgreSQL credentials
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: postgres-credentials
  namespace: postgresql
type: Opaque
stringData:
  POSTGRES_DB: ps_db
  POSTGRES_USER: ps_user
  POSTGRES_PASSWORD: ${PASSWORD}
EOF

# Create ConfigMap for pgBackRest configuration
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: pgbackrest-config
  namespace: postgresql
data:
  pgbackrest.conf: |
    [global]
    repo1-path=/var/lib/pgbackrest
    log-level-console=info
    log-level-file=detail
    start-fast=y
    process-max=4

    [main]
    pg1-path=/var/lib/postgresql/17/main
EOF

# Create ConfigMap for PostgreSQL initialization
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: postgres-initdb
  namespace: postgresql
data:
  init-extensions.sh: |
    #!/bin/bash
    set -e
    
    psql -v ON_ERROR_STOP=1 --username "\$POSTGRES_USER" --dbname "\$POSTGRES_DB" <<-EOSQL
      CREATE EXTENSION IF NOT EXISTS hstore;
      CREATE EXTENSION IF NOT EXISTS postgis;
      CREATE EXTENSION IF NOT EXISTS pg_stat_statements;
      CREATE EXTENSION IF NOT EXISTS pgcrypto;
      CREATE EXTENSION IF NOT EXISTS pg_partman;
      CREATE EXTENSION IF NOT EXISTS postgres_fdw;
      CREATE EXTENSION IF NOT EXISTS vector;
      CREATE EXTENSION IF NOT EXISTS pgpcre;
      CREATE EXTENSION IF NOT EXISTS timescaledb_toolkit;
    EOSQL
EOF
