# PostgreSQL 17 with TimescaleDB 2.19.3

This Docker image provides PostgreSQL 17 with TimescaleDB 2.19.3 and additional extensions optimized for production use.

## Features

- PostgreSQL 17
- TimescaleDB 2.19.3
- pg_partman for automated table partitioning
- pgBackRest for backup and restore
- PostGIS 3 for spatial data
- pgvector for vector operations
- timescaledb-toolkit for advanced time-series analytics
- pgpcre for Perl-compatible regular expressions

## Quick Start

```bash
# Pull the image
docker pull yourusername/pg17-4-timescaledb:2.19.3-slim

# Run a container
docker run -d --name timescaledb \
  -p 5432:5432 \
  -e POSTGRES_PASSWORD=yourpassword \
  -v pg_data:/var/lib/postgresql/17/main \
  -v pgbackrest_conf:/etc/pgbackrest \
  -v pgbackrest_data:/var/lib/pgbackrest \
  -v pgbackrest_logs:/var/log/pgbackrest \
  yourusername/pg17-4-timescaledb:2.19.3-slim
```

Or use the provided docker-compose.yml file:

```bash
docker-compose up -d
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `POSTGRES_PASSWORD` | PostgreSQL superuser password | (required) |
| `PGBACKREST_LOG_LEVEL_CONSOLE` | pgBackRest console log level | info |
| `PGBACKREST_LOG_LEVEL_FILE` | pgBackRest file log level | detail |
| `PGBACKREST_START_FAST` | pgBackRest start fast option | y |
| `PGBACKREST_PROCESS_MAX` | pgBackRest maximum parallel processes | 4 |
| `PGBACKREST_REPO1_PATH` | pgBackRest repository path | /var/lib/pgbackrest |
| `PGBACKREST_LOG_PATH` | pgBackRest log path | (optional) |

## Backup Configuration

This image comes with pgBackRest pre-configured. To customize the backup settings, mount a volume to `/etc/pgbackrest` and provide your own pgbackrest.conf file.

Basic backup example:

```bash
docker exec -it timescaledb gosu postgres pgbackrest --stanza=main backup
```

## Volumes

- `/var/lib/postgresql/17/main`: PostgreSQL data directory
- `/etc/pgbackrest`: pgBackRest configuration
- `/var/lib/pgbackrest`: pgBackRest repository
- `/var/log/pgbackrest`: pgBackRest logs

## Extensions

To enable the extensions in your database:

```sql
CREATE EXTENSION timescaledb;
CREATE EXTENSION postgis;
CREATE EXTENSION pg_partman;
CREATE EXTENSION pgvector;
CREATE EXTENSION timescaledb_toolkit;
CREATE EXTENSION pgpcre;
```

## Building the Image

```bash
# Clone this repository
git clone https://github.com/yourusername/postgres-timescaledb.git
cd postgres-timescaledb

# Build the image
chmod +x build-push.sh
./build-push.sh
```

## Security Recommendations

- Always use strong passwords
- Configure proper network access controls
- Consider using PostgreSQL SSL connections
- Review and customize pg_hba.conf settings as needed
- Run regular security updates

## License

This Dockerfile and associated scripts are provided under the MIT License. TimescaleDB and PostgreSQL have their own licenses.

## Credits

This image bundles software from the following projects:

- [PostgreSQL](https://postgresql.org/)
- [TimescaleDB](https://www.timescale.com/)
- [pgBackRest](https://pgbackrest.org/)
- [PostGIS](https://postgis.net/)
- [pg_partman](https://github.com/pgpartman/pg_partman)
- [pgvector](https://github.com/pgvector/pgvector)
- [pgpcre](https://github.com/petere/pgpcre)
