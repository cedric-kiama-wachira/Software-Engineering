# Multi-stage build to reduce final image size

# Builder stage for Go components
FROM golang:1.24-bullseye AS go_builder

# Build timescaledb-tune
RUN cd /tmp && \
    wget -q https://github.com/timescale/timescaledb-tune/archive/refs/tags/v0.18.0.tar.gz && \
    tar zxf v0.18.0.tar.gz && \
    cd timescaledb-tune-0.18.0/ && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o timescaledb-tune ./cmd/...

# Builder stage for TimescaleDB
FROM debian:bookworm-slim AS timescaledb_builder

# Install basic build dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    build-essential \
    ca-certificates \
    cmake \
    curl \
    git \
    libssl-dev \
    && rm -rf /var/lib/apt/lists/*

# Add PostgreSQL repository
RUN mkdir -p /usr/share/postgresql-common/pgdg && \
    curl -sSL -o /usr/share/postgresql-common/pgdg/apt.postgresql.org.asc \
    https://www.postgresql.org/media/keys/ACCC4CF8.asc && \
    echo 'deb [signed-by=/usr/share/postgresql-common/pgdg/apt.postgresql.org.asc] https://apt.postgresql.org/pub/repos/apt bookworm-pgdg main' \
    > /etc/apt/sources.list.d/pgdg.list && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
    postgresql-server-dev-17 \
    && rm -rf /var/lib/apt/lists/*

# Build TimescaleDB
RUN git clone --depth 1 --branch 2.19.3 https://github.com/timescale/timescaledb.git && \
    cd timescaledb && \
    mkdir build && \
    cd build && \
    cmake .. -DREGRESS_CHECKS=OFF -DCMAKE_BUILD_TYPE=Release && \
    make -j$(nproc) && \
    make install

# Builder stage for pg_partman
FROM debian:bookworm-slim AS pg_partman_builder

# Install build dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    build-essential \
    ca-certificates \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*

# Add PostgreSQL repository
RUN mkdir -p /usr/share/postgresql-common/pgdg && \
    curl -sSL -o /usr/share/postgresql-common/pgdg/apt.postgresql.org.asc \
    https://www.postgresql.org/media/keys/ACCC4CF8.asc && \
    echo 'deb [signed-by=/usr/share/postgresql-common/pgdg/apt.postgresql.org.asc] https://apt.postgresql.org/pub/repos/apt bookworm-pgdg main' \
    > /etc/apt/sources.list.d/pgdg.list && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
    postgresql-server-dev-17 \
    && rm -rf /var/lib/apt/lists/*

# Build pg_partman
RUN git clone --depth 1 https://github.com/pgpartman/pg_partman.git && \
    cd pg_partman && \
    make -j$(nproc) && \
    make install

# Builder stage for pgpcre
FROM debian:bookworm-slim AS pgpcre_builder

# Install build dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    build-essential \
    ca-certificates \
    curl \
    git \
    libpcre3-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

# Add PostgreSQL repository
RUN mkdir -p /usr/share/postgresql-common/pgdg && \
    curl -sSL -o /usr/share/postgresql-common/pgdg/apt.postgresql.org.asc \
    https://www.postgresql.org/media/keys/ACCC4CF8.asc && \
    echo 'deb [signed-by=/usr/share/postgresql-common/pgdg/apt.postgresql.org.asc] https://apt.postgresql.org/pub/repos/apt bookworm-pgdg main' \
    > /etc/apt/sources.list.d/pgdg.list && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
    postgresql-server-dev-17 \
    && rm -rf /var/lib/apt/lists/*

# Build pgpcre with proper linking to PCRE library
RUN git clone --depth 1 https://github.com/petere/pgpcre.git && \
    cd pgpcre && \
    # Make sure the Makefile properly links against PCRE library
    echo 'SHLIB_LINK = -lpcre' >> Makefile && \
    make -j$(nproc) && \
    make install

# Builder stage for Rust components
FROM debian:bookworm-slim AS rust_builder

# Install build dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    build-essential \
    ca-certificates \
    curl \
    git \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

# Add PostgreSQL repository
RUN mkdir -p /usr/share/postgresql-common/pgdg && \
    curl -sSL -o /usr/share/postgresql-common/pgdg/apt.postgresql.org.asc \
    https://www.postgresql.org/media/keys/ACCC4CF8.asc && \
    echo 'deb [signed-by=/usr/share/postgresql-common/pgdg/apt.postgresql.org.asc] https://apt.postgresql.org/pub/repos/apt bookworm-pgdg main' \
    > /etc/apt/sources.list.d/pgdg.list && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
    postgresql-server-dev-17 \
    libclang-14-dev \
    && rm -rf /var/lib/apt/lists/*

# Install Rust
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y && \
    . "$HOME/.cargo/env" && \
    rustup default stable

# Build timescaledb-toolkit
RUN . "$HOME/.cargo/env" && \
    cargo install --version '=0.12.9' --locked cargo-pgrx && \
    cargo pgrx init --pg17 $(which pg_config) && \
    git clone --depth 1 https://github.com/timescale/timescaledb-toolkit && \
    cd timescaledb-toolkit/extension && \
    cargo pgrx install --release

# Final stage
FROM debian:bookworm-slim

# Add PostgreSQL repository
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /usr/share/postgresql-common/pgdg && \
    curl -sSL -o /usr/share/postgresql-common/pgdg/apt.postgresql.org.asc \
    https://www.postgresql.org/media/keys/ACCC4CF8.asc && \
    echo 'deb [signed-by=/usr/share/postgresql-common/pgdg/apt.postgresql.org.asc] https://apt.postgresql.org/pub/repos/apt bookworm-pgdg main' \
    > /etc/apt/sources.list.d/pgdg.list

# Install runtime dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    gosu \
    libpcre3 \
    libssl3 \
    openssl \
    postgresql-17 \
    postgresql-client-17 \
    postgresql-17-postgis-3 \
    postgresql-17-pgvector \
    pgbackrest \
    && rm -rf /var/lib/apt/lists/*

# Copy components from builders
COPY --from=go_builder /tmp/timescaledb-tune-0.18.0/timescaledb-tune /usr/local/bin/

# TimescaleDB files
COPY --from=timescaledb_builder /usr/lib/postgresql/17/lib/timescaledb*.so /usr/lib/postgresql/17/lib/
COPY --from=timescaledb_builder /usr/share/postgresql/17/extension/timescaledb* /usr/share/postgresql/17/extension/

# pg_partman files
COPY --from=pg_partman_builder /usr/lib/postgresql/17/lib/pg_partman_bgw.so /usr/lib/postgresql/17/lib/
COPY --from=pg_partman_builder /usr/share/postgresql/17/extension/pg_partman* /usr/share/postgresql/17/extension/

# pgpcre files
COPY --from=pgpcre_builder /usr/lib/postgresql/17/lib/pgpcre.so /usr/lib/postgresql/17/lib/
COPY --from=pgpcre_builder /usr/share/postgresql/17/extension/pgpcre* /usr/share/postgresql/17/extension/

# timescaledb-toolkit files
COPY --from=rust_builder /usr/lib/postgresql/17/lib/timescaledb_toolkit*.so /usr/lib/postgresql/17/lib/
COPY --from=rust_builder /usr/share/postgresql/17/extension/timescaledb_toolkit* /usr/share/postgresql/17/extension/

# Configure pgBackRest
RUN mkdir -p /var/log/pgbackrest && \
    chown postgres:postgres /var/log/pgbackrest && \
    mkdir -p /etc/pgbackrest/conf.d && \
    touch /etc/pgbackrest/pgbackrest.conf && \
    chmod 640 /etc/pgbackrest/pgbackrest.conf && \
    chown postgres:postgres /etc/pgbackrest/pgbackrest.conf

# Clean up to reduce image size
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /tmp/* /var/tmp/* && \
    find /var/log -type f -delete && \
    rm -rf /usr/share/doc /usr/share/man /usr/share/info

# Entrypoint
COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

EXPOSE 5432
VOLUME /var/lib/postgresql/17/main
VOLUME /etc/pgbackrest
VOLUME /var/lib/pgbackrest
VOLUME /var/log/pgbackrest

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["postgres"]
