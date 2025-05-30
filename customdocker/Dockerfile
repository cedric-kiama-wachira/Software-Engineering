# Multi-stage build to reduce final image size

# ARGs for build context
ARG BUILD_DATE
ARG VCS_REF
ARG VERSION=2.19.3-slim

# Builder stage for Go components - Using security best practices
FROM golang:1.24-bullseye AS go_builder

# Create non-root user for build
RUN useradd -r -u 10001 -g 0 builder && \
    mkdir -p /tmp/build && \
    chown -R builder:0 /tmp/build && \
    chmod -R 755 /tmp/build

USER builder
WORKDIR /tmp/build

# Configure Go to use /tmp for cache to avoid permission issues
ENV GOCACHE=/tmp/gocache 
ENV GOPATH=/tmp/go
ENV PATH=$PATH:/tmp/go/bin

# Build timescaledb-tune
RUN mkdir -p /tmp/gocache /tmp/go && \
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
    make install && \
    cd / && rm -rf /timescaledb && apt-get remove -y git && apt-get autoremove -y

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
    make install && \
    cd / && rm -rf /pg_partman && apt-get remove -y git && apt-get autoremove -y

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
    echo 'SHLIB_LINK = -lpcre' >> Makefile && \
    make -j$(nproc) && \
    make install && \
    cd / && rm -rf /pgpcre && apt-get remove -y git && apt-get autoremove -y

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
    postgresql-17 \
    postgresql-server-dev-17 \
    libclang-14-dev \
    && rm -rf /var/lib/apt/lists/*

# Create user for Rust builds with proper permissions
RUN useradd -m -s /bin/bash rustuser && \
    mkdir -p /tmp/rust-build && \
    chown -R rustuser:rustuser /tmp/rust-build && \
    # Give rustuser write permissions to PostgreSQL extension directories
    mkdir -p /usr/share/postgresql/17/extension && \
    mkdir -p /usr/lib/postgresql/17/lib && \
    chown -R rustuser:rustuser /usr/share/postgresql/17/extension && \
    chown -R rustuser:rustuser /usr/lib/postgresql/17/lib

USER rustuser
WORKDIR /tmp/rust-build

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

USER root
RUN rm -rf /timescaledb-toolkit && apt-get remove -y git && apt-get autoremove -y

# Final stage
FROM debian:bookworm-slim

# Install required packages including curl for later use
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    gnupg \
    lsb-release \
    && rm -rf /var/lib/apt/lists/*

# Now add PostgreSQL repository (curl is available now)
RUN mkdir -p /usr/share/postgresql-common/pgdg && \
    curl -sSL -o /usr/share/postgresql-common/pgdg/apt.postgresql.org.asc \
    https://www.postgresql.org/media/keys/ACCC4CF8.asc && \
    echo 'deb [signed-by=/usr/share/postgresql-common/pgdg/apt.postgresql.org.asc] https://apt.postgresql.org/pub/repos/apt bookworm-pgdg main' \
    > /etc/apt/sources.list.d/pgdg.list

# Install runtime dependencies without version pinning
# Replace gosu with util-linux (for setpriv)
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    util-linux \
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
COPY --link --from=go_builder /tmp/build/timescaledb-tune-0.18.0/timescaledb-tune /usr/local/bin/

# TimescaleDB files
COPY --link --from=timescaledb_builder /usr/lib/postgresql/17/lib/timescaledb*.so /usr/lib/postgresql/17/lib/
COPY --link --from=timescaledb_builder /usr/share/postgresql/17/extension/timescaledb* /usr/share/postgresql/17/extension/

# pg_partman files
COPY --link --from=pg_partman_builder /usr/lib/postgresql/17/lib/pg_partman_bgw.so /usr/lib/postgresql/17/lib/
COPY --link --from=pg_partman_builder /usr/share/postgresql/17/extension/pg_partman* /usr/share/postgresql/17/extension/

# pgpcre files
COPY --link --from=pgpcre_builder /usr/lib/postgresql/17/lib/pgpcre.so /usr/lib/postgresql/17/lib/
COPY --link --from=pgpcre_builder /usr/share/postgresql/17/extension/pgpcre* /usr/share/postgresql/17/extension/

# timescaledb-toolkit files
COPY --link --from=rust_builder /usr/lib/postgresql/17/lib/timescaledb_toolkit*.so /usr/lib/postgresql/17/lib/
COPY --link --from=rust_builder /usr/share/postgresql/17/extension/timescaledb_toolkit* /usr/share/postgresql/17/extension/

# Configure pgBackRest with proper permissions
RUN mkdir -p /var/log/pgbackrest && \
    chown postgres:postgres /var/log/pgbackrest && \
    chmod 750 /var/log/pgbackrest && \
    mkdir -p /etc/pgbackrest/conf.d && \
    touch /etc/pgbackrest/pgbackrest.conf && \
    chmod 640 /etc/pgbackrest/pgbackrest.conf && \
    chown postgres:postgres /etc/pgbackrest/pgbackrest.conf

# Clean up to reduce image size and attack surface
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /tmp/* /var/tmp/* && \
    find /var/log -type f -delete && \
    rm -rf /usr/share/doc /usr/share/man /usr/share/info

# Set appropriate file permissions for PostgreSQL directories
RUN mkdir -p /var/lib/postgresql/17/main && \
    chown -R postgres:postgres /var/lib/postgresql && \
    chmod -R 700 /var/lib/postgresql/17/main

# Copy entrypoint script
COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

# Add OCI labels for better traceability
LABEL org.opencontainers.image.title="PostgreSQL with TimescaleDB"
LABEL org.opencontainers.image.description="PostgreSQL 17 with TimescaleDB, PostGIS, pgVector and other extensions"
LABEL org.opencontainers.image.version="${VERSION}"
LABEL org.opencontainers.image.created="${BUILD_DATE}"
LABEL org.opencontainers.image.revision="${VCS_REF}"
LABEL org.opencontainers.image.vendor="Cedric Kiama"
LABEL org.opencontainers.image.licenses="PostgreSQL"
LABEL org.opencontainers.image.source="https://github.com/cedrickiama/pg17-4-timescaledb"
LABEL org.opencontainers.image.base.name="debian:bookworm-slim"

# Define volumes and expose port
EXPOSE 5432
VOLUME /var/lib/postgresql/17/main
VOLUME /etc/pgbackrest
VOLUME /var/lib/pgbackrest
VOLUME /var/log/pgbackrest

# Add healthcheck
HEALTHCHECK --interval=30s --timeout=10s --start-period=40s --retries=3 \
  CMD pg_isready -U postgres || exit 1

USER root
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["postgres"]
