#!/bin/bash
set -e

# Validate critical environment variables
if [ -z "$POSTGRES_PASSWORD" ]; then
    echo "[$(date +'%Y-%m-%d %H:%M:%S')] ERROR: POSTGRES_PASSWORD environment variable is not set."
    exit 1
fi

if [ -z "$PGBACKREST_CIPHER_PASS" ]; then
    echo "[$(date +'%Y-%m-%d %H:%M:%S')] ERROR: PGBACKREST_CIPHER_PASS environment variable is not set."
    exit 1
fi

# Function to print message with timestamp
log() {
    echo "[$(date +'%Y-%m-%d %H:%M:%S')] $1"
}

# Security enhancement: Verify file permissions before starting
check_permissions() {
    log "Verifying file permissions..."
    
    # Check pgbackrest directories
    if [ -d /var/log/pgbackrest ]; then
        chmod 750 /var/log/pgbackrest
        chown -R postgres:postgres /var/log/pgbackrest
    fi
    
    if [ -d /etc/pgbackrest ]; then
        chmod 750 /etc/pgbackrest
        chown -R postgres:postgres /etc/pgbackrest
    fi
    
    if [ -d /var/lib/pgbackrest ]; then
        chmod 750 /var/lib/pgbackrest
        chown -R postgres:postgres /var/lib/pgbackrest
    fi
    
    # Check PostgreSQL data directory
    if [ -d /var/lib/postgresql/17/main ]; then
        chmod 700 /var/lib/postgresql/17/main
        chown -R postgres:postgres /var/lib/postgresql/17/main
    fi
    
    # Check entrypoint script
    chmod 755 /usr/local/bin/entrypoint.sh
}

# Configure pgBackRest with environment variables and encryption
configure_pgbackrest() {
    # Create base config if it doesn't exist yet or is empty
    if [ ! -s /etc/pgbackrest/pgbackrest.conf ]; then
        log "Setting up default pgbackrest.conf..."
        cat > /etc/pgbackrest/pgbackrest.conf << EOF
[global]
repo1-path=/var/lib/pgbackrest
log-level-console=${PGBACKREST_LOG_LEVEL_CONSOLE:-info}
log-level-file=${PGBACKREST_LOG_LEVEL_FILE:-detail}
start-fast=${PGBACKREST_START_FAST:-y}
process-max=${PGBACKREST_PROCESS_MAX:-4}
log-timestamp=y
EOF

        # Add encryption if cipher pass is available
        if [ -n "$PGBACKREST_CIPHER_PASS" ]; then
            log "Configuring backup encryption..."
            cat >> /etc/pgbackrest/pgbackrest.conf << EOF
repo1-cipher-type=aes-256-cbc
compress-type=zst
compress-level=6
EOF
        fi

        cat >> /etc/pgbackrest/pgbackrest.conf << EOF
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

    # Secure permissions
    chmod 640 /etc/pgbackrest/pgbackrest.conf
}

# Function to update PostgreSQL password with improved security
update_postgres_password() {
    if [ -n "$POSTGRES_PASSWORD" ]; then
        log "Setting postgres password..."
        # Security enhancement: Check password strength
        if [ ${#POSTGRES_PASSWORD} -lt 8 ]; then
            log "WARNING: Password is less than 8 characters. Consider using a stronger password."
        fi
        
        su postgres -c "psql -c \"ALTER USER postgres WITH PASSWORD '${POSTGRES_PASSWORD}';\""
        
        # Update pg_hba.conf to use SCRAM-SHA-256 authentication (more secure than MD5)
        log "Updating authentication settings with SCRAM-SHA-256..."
        cat > /etc/postgresql/17/main/pg_hba.conf << EOF
# TYPE  DATABASE        USER            ADDRESS                 METHOD
local   all             postgres                                scram-sha-256
host    all             all             127.0.0.1/32            scram-sha-256
host    all             all             ::1/128                 scram-sha-256
host    all             all             0.0.0.0/0               scram-sha-256
local   all             all                                     trust
EOF
        
        # Security enhancement: Update postgresql.conf with secure settings
        cat >> /etc/postgresql/17/main/postgresql.conf << EOF
# Security settings
password_encryption = 'scram-sha-256'  # More secure password hashing
ssl = on                               # Enable SSL
ssl_prefer_server_ciphers = on         # Server ciphers preference
EOF
        
        # Reload PostgreSQL to apply new settings
        pg_ctlcluster 17 main reload
    fi
}

# Check if running as root (Docker can be configured to run as non-root)
if [ "$(id -u)" = "0" ]; then
    # Perform security checks and setup
    check_permissions
    
    # Check if data directory exists and is empty
    if [ -z "$(ls -A /var/lib/postgresql/17/main 2>/dev/null)" ]; then
        log "Initializing PostgreSQL cluster..."
        pg_createcluster 17 main --start

        log "Configuring initial authentication..."
        echo "local all all trust" > /etc/postgresql/17/main/pg_hba.conf
        echo "host all all 0.0.0.0/0 trust" >> /etc/postgresql/17/main/pg_hba.conf
        echo "host all all ::1/128 trust" >> /etc/postgresql/17/main/pg_hba.conf
        echo "host all all 127.0.0.1/32 trust" >> /etc/postgresql/17/main/pg_hba.conf
        
        log "Configuring PostgreSQL settings..."
        echo "listen_addresses = '*'" >> /etc/postgresql/17/main/postgresql.conf
        echo "shared_preload_libraries = 'timescaledb,pg_partman_bgw'" >> /etc/postgresql/17/main/postgresql.conf
        echo "pg_partman_bgw.interval = 3600" >> /etc/postgresql/17/main/postgresql.conf

        # Critical: Restart to load shared libraries BEFORE creating extensions
        log "Restarting PostgreSQL to apply configuration..."
        pg_ctlcluster 17 main restart
        sleep 2

        update_postgres_password
        
        log "Waiting for PostgreSQL to become ready..."
        export PGPASSWORD="$POSTGRES_PASSWORD"
        until su postgres -c "psql -h localhost -U postgres -c '\q'"
        do
            log "Waiting for PostgreSQL..."
            sleep 2
        done
        unset PGPASSWORD

        log "Configuring TimescaleDB..."
        timescaledb-tune --quiet --yes
        
        # Configure pgbackrest
        configure_pgbackrest
        
        pg_ctlcluster 17 main stop
    else
        if [ -f /var/lib/postgresql/17/main/PG_VERSION ]; then
            log "Starting PostgreSQL temporarily to update configuration..."
            pg_ctlcluster 17 main start
            sleep 3
            update_postgres_password
            
            # Make sure pgbackrest is configured even on restart
            configure_pgbackrest
            
            pg_ctlcluster 17 main stop
        fi
    fi

    # Final configuration checks
    log "Ensuring proper configuration..."
    if [ ! -f /etc/postgresql/17/main/postgresql.conf.timescaledb-tune.backup ]; then
        log "Running timescaledb-tune for optimal performance..."
        timescaledb-tune --quiet --yes || true
    fi

    log "Starting PostgreSQL as postgres user..."
exec setpriv --reuid=$(id -u postgres) --regid=$(id -g postgres) --init-groups \
    /usr/lib/postgresql/17/bin/postgres \
    -D /var/lib/postgresql/17/main \
    -c config_file=/etc/postgresql/17/main/postgresql.conf

else
    # Already running as postgres user
    log "Starting PostgreSQL..."
    exec /usr/lib/postgresql/17/bin/postgres \
        -D /var/lib/postgresql/17/main \
        -c config_file=/etc/postgresql/17/main/postgresql.conf
fi
