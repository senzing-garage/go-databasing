# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

go-databasing is a Go library providing direct database access for Senzing databases. This is a Senzing Garage project (experimental/development) with semantic versions `0.n.x` - not production-ready.

**Supported databases:** PostgreSQL, SQLite, MySQL, MS SQL Server, Oracle, DB2 (partial)

## Build Commands

```bash
# First-time setup (installs golangci-lint, gotestfmt, govulncheck, godoc, etc.)
make dependencies-for-development

# Update Go module dependencies
make dependencies

# Build for current platform
make build

# Build for all platforms (darwin/linux/windows × amd64/arm64)
make build-all
```

## Testing

Tests require running database services via Docker Compose.

```bash
# Start database services (PostgreSQL, MySQL, MS SQL, Oracle)
make setup

# Run tests (with gotestfmt output)
make test

# Run tests with coverage report (opens HTML in browser)
make coverage

# Verify coverage meets thresholds
make check-coverage

# Stop services and clean up
make clean
```

Run a single test:
```bash
go test -v -run TestFunctionName ./path/to/package
```

Tests run with `-p 1` (serial execution) due to database resource constraints.

## Linting

```bash
# Run all linters (golangci-lint + govulncheck + cspell)
make lint

# Individual linters
make golangci-lint
make govulncheck
make cspell
make bearer          # Security scanning

# Auto-fix many lint issues
make fix
```

The project uses an extensive golangci-lint configuration (`.github/linters/.golangci.yaml`) with 80+ linters enabled. Code coverage threshold is >80%.

## Architecture

### Package Structure

- **connector/** - Factory for creating `database/sql/driver.Connector` from database URLs
  - `NewConnector(ctx, databaseURL)` routes to database-specific connectors based on URL scheme
- **connector{mysql,postgresql,sqlite,mssql,oracle,db2}/** - Database-specific connector implementations
- **sqlexecutor/** - Reads SQL files and executes statements against databases
  - `BasicSQLExecutor` implements the `SQLExecutor` interface
  - Uses Senzing component ID 6422 for logging
- **postgresql/** - PostgreSQL-specific utilities (e.g., `GetCurrentWatermark()` for transaction OID)
- **checker/** - Validates Senzing database schema installation
- **dbhelper/** - OS-specific database helpers

### Database URL Schemes

```
sqlite3://na:na@nowhere/path/to/db.db
postgresql://user:pass@host:5432/database
mysql://user:pass@host:3306/database
mssql://user:pass@host:1433/database
oci://user:pass@host:1521/service
```

### Key Patterns

- All packages follow the Senzing `Basic*` struct pattern (e.g., `BasicSQLExecutor`, `BasicPostgresql`)
- Observer pattern support via `go-observing` for monitoring database operations
- Structured logging via `go-logging` and `go-messaging`

## Environment Variables

- `LD_LIBRARY_PATH` - Oracle instant client and Senzing libraries (default: `/opt/senzing/er/lib:/opt/oracle/instantclient_23_5`)
- `SENZING_TOOLS_DATABASE_URL` - Default database URL for testing (default: `sqlite3://na:na@nowhere/tmp/sqlite/G2C.db`)
- `SENZING_LOG_LEVEL` - Set to `TRACE` for verbose logging during coverage tests

## Test Database Access (via Docker)

When running `make setup`:
- PostgreSQL: localhost:5432 (pgAdmin at localhost:9171)
- MySQL: localhost:3306 (phpMyAdmin at localhost:9173)
- MS SQL: localhost:1433 (Adminer at localhost:9177)
- Oracle: localhost:1521
