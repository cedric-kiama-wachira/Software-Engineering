kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: postgres-init-script
  namespace: postgresql
data:
  init-extensions.sql: |
    CREATE EXTENSION timescaledb;
    CREATE EXTENSION pgpcre;
    CREATE EXTENSION EXISTS postgis;
    CREATE EXTENSION EXISTS hstore;
    CREATE EXTENSION pg_stat_statements;
    CREATE EXTENSION pgcrypto;
    CREATE EXTENSION pg_partman;
    CREATE EXTENSION postgres_fdw;
    CREATE EXTENSION vector;
    CREATE EXTENSION timescaledb_toolkit;
EOF
