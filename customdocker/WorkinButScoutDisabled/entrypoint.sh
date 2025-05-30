#!/bin/bash
set -e

# Configure pgBackRest with environment variables if provided
configure_pgbackrest() {
    # Create base config if it doesn't exist yet or is empty
    if [ ! -s /etc/pgbackrest/pgbackrest.conf ]; then
        echo "Setting up default pgbackrest.conf..."
        cat > /etc/pgbackrest/pgbackrest.conf << EOF
[global]
repo1-path=/var/lib/pgbackrest
log-level-console=${PGBACKREST_LOG_LEVEL_CONSOLE:-info}
log-level-file=${PGBACKREST_LOG_LEVEL_FILE:-detail}
start-fast=${PGBACKREST_START_FAST:-y}
process-max=${PGBACKREST_PROCESS_MAX:-4}

[main]
pg1-path=/var/lib/postgresql/17/main
EOF
    fi

    # Apply any custom environment variables to the pgbackrest.conf
    if [ -n "$PGBACKREST_REPO1_PATH" ]; then
        sed -i "s|repo1-path=.*|repo1-path=${PGBACKREST_REPO1_PATH}|g" /etc/pgbackrest/pgbackrest.conf
    fi
    
    if [ -n "$PGBACKREST_LOG_PATH" ]; then
        echo "log-path=${PGBACKREST_LOG_PATH}" >> /etc/pgbackrest/pgbackrest.conf
    fi
    
    # Check ownership on directories
    chown -R postgres:postgres /var/log/pgbackrest
    chown -R postgres:postgres /etc/pgbackrest
    chown -R postgres:postgres /var/lib/pgbackrest 2>/dev/null || true
}

# Function to update PostgreSQL password
update_postgres_password() {
    if [ -n "$POSTGRES_PASSWORD" ]; then
        echo "Setting postgres password..."
        su postgres -c "psql -c \"ALTER USER postgres WITH PASSWORD '${POSTGRES_PASSWORD}';\""
        
        # Update pg_hba.conf to use the password
        echo "Updating authentication settings..."
        echo "local all postgres scram-sha-256" > /etc/postgresql/17/main/pg_hba.conf
        echo "host all all 0.0.0.0/0 scram-sha-256" >> /etc/postgresql/17/main/pg_hba.conf
        echo "host all all ::1/128 scram-sha-256" >> /etc/postgresql/17/main/pg_hba.conf
        echo "host all all 127.0.0.1/32 scram-sha-256" >> /etc/postgresql/17/main/pg_hba.conf
        echo "local all all trust" >> /etc/postgresql/17/main/pg_hba.conf
        
        # Reload PostgreSQL to apply new settings
        pg_ctlcluster 17 main reload
    fi
}

# Check if data directory exists and is empty
if [ -z "$(ls -A /var/lib/postgresql/17/main 2>/dev/null)" ]; then
    echo "Initializing PostgreSQL cluster..."
    pg_createcluster 17 main --start

    echo "Configuring initial authentication..."
    echo "local all all trust" > /etc/postgresql/17/main/pg_hba.conf
    echo "host all all 0.0.0.0/0 trust" >> /etc/postgresql/17/main/pg_hba.conf
    echo "host all all ::1/128 trust" >> /etc/postgresql/17/main/pg_hba.conf
    echo "host all all 127.0.0.1/32 trust" >> /etc/postgresql/17/main/pg_hba.conf
    
    echo "Configuring PostgreSQL settings..."
    echo "listen_addresses = '*'" >> /etc/postgresql/17/main/postgresql.conf
    echo "shared_preload_libraries = 'timescaledb,pg_partman_bgw'" >> /etc/postgresql/17/main/postgresql.conf
    echo "pg_partman_bgw.interval = 3600" >> /etc/postgresql/17/main/postgresql.conf

    # Critical: Restart to load shared libraries BEFORE creating extensions
    echo "Restarting PostgreSQL to apply configuration..."
    pg_ctlcluster 17 main restart
    sleep 2

    update_postgres_password
    
    echo "Waiting for PostgreSQL to become ready..."
    export PGPASSWORD="$POSTGRES_PASSWORD"
    until su postgres -c "psql -h localhost -U postgres -c '\q'"
    do
        echo "Waiting for PostgreSQL..."
        sleep 2
    done
    unset PGPASSWORD

    echo "Configuring TimescaleDB..."
    timescaledb-tune --quiet --yes
    
    # Configure pgbackrest
    configure_pgbackrest
    
    pg_ctlcluster 17 main stop
else
    if [ -f /var/lib/postgresql/17/main/PG_VERSION ]; then
        echo "Starting PostgreSQL temporarily to update configuration..."
        pg_ctlcluster 17 main start
        sleep 3
        update_postgres_password
        
        # Make sure pgbackrest is configured even on restart
        configure_pgbackrest
        
        pg_ctlcluster 17 main stop
    fi
fi

# Final configuration checks
echo "Ensuring proper configuration..."
if [ ! -f /etc/postgresql/17/main/postgresql.conf.timescaledb-tune.backup ]; then
    echo "Running timescaledb-tune for optimal performance..."
    timescaledb-tune --quiet --yes || true
fi

echo "Starting PostgreSQL as postgres user..."
exec gosu postgres /usr/lib/postgresql/17/bin/postgres \
    -D /var/lib/postgresql/17/main \
    -c config_file=/etc/postgresql/17/main/postgresql.conf
