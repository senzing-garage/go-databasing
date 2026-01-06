# Changelog

All notable changes to this project will be documented in this file.

The changelog format is based on [Keep a Changelog] and [CommonMark].
This project adheres to [Semantic Versioning].

## [Unreleased]

-

## [0.5.9] - 2026-01-06

### Changed in 0.5.9

- Update dependencies
- Improved string building performance in connector package

## [0.5.8] - 2025-05-20

### Changed in 0.5.8

- Improved error messages

## [0.5.7] - 2025-04-25

### Changed in 0.5.7

- Update dependencies

## [0.5.6] - 2025-03-13

### Changed in 0.5.6

- Strip trailing semi-colons in SQL files before processing

## [0.5.5] - 2025-03-12

### Changed in 0.5.5

- Oracle database protocol change from `oracle` to `oci`
- Update dependencies

## [0.5.4] - 2024-11-14

### Added in 0.5.4

- Support for in-memory SQLite

## [0.5.3] - 2024-10-18

### Changed in 0.5.3

- Update dependencies

## [0.5.2] - 2024-09-09

### Changed in 0.5.2

- Support for Oracle database

## [0.5.1] - 2024-09-05

### Changed in 0.5.1

- Update dependencies
- Improve documentation

## [0.5.0] - 2024-08-20

### Changed in 0.5.0

- Change from `g2` to `sz`/`er`

## [0.4.2] - 2024-06-11

### Changed in 0.4.2

- From `CheckerImpl` to `BasicChecker`
- From `PostgresImpl` to `BasicPostgresql`
- From `SqlExecutorImpl` to `BasicSQLExecutor`

## [0.4.1] - 2024-04-19

### Changed in 0.4.1

- Update dependencies
  - github.com/go-sql-driver/mysql v1.8.1
  - github.com/mattn/go-sqlite3 v1.14.22
  - github.com/microsoft/go-mssqldb v1.7.1
  - github.com/stretchr/testify v1.9.0

## [0.4.0] - 2024-01-02

### Changed in 0.4.0

- Renamed module to `github.com/senzing-garage/go-databasing`
- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/mattn/go-sqlite3 v1.14.19
  - github.com/senzing-garage/go-logging v1.4.0
  - github.com/senzing-garage/go-observing v0.3.0

## [0.3.1] - 2023-10-17

### Changed in 0.3.1

- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/senzing-garage/go-logging v1.3.3
  - github.com/senzing-garage/go-observing v0.2.8

## [0.3.0] - 2023-10-06

### Added to 0.3.0

- `checker.IsSchemaInstalled`

### Changed in 0.3.0

- Updated `.sql` files

## [0.2.9] - 2023-08-31

### Added in 0.2.9

- `dbhelper.ExtractSqliteDatabaseFilename()`

## [0.2.8] - 2023-08-30

### Changed in 0.2.8

- Moved examples to separate file

## [0.2.7] - 2023-08-04

### Changed in 0.2.7

- Refactor to `template-go`
- Updated dependencies
  - github.com/microsoft/go-mssqldb v1.5.0
  - github.com/senzing-garage/go-logging v1.3.2
  - github.com/senzing-garage/go-observing v0.2.7

## [0.2.6] - 2023-07-13

### Added in 0.2.6

- `SetObserverOrigin()` methods

### Changed in 0.2.6

- Updated dependencies
  - github.com/microsoft/go-mssqldb v1.3.0
  - github.com/senzing-garage/go-logging v1.3.1

## [0.2.5] - 2023-06-16

### Changed in 0.2.5

- Updated dependencies
  - github.com/mattn/go-sqlite3 v1.14.17
  - github.com/microsoft/go-mssqldb v1.1.0
  - github.com/senzing-garage/go-logging v1.2.6
  - github.com/senzing-garage/go-observing v0.2.6
  - github.com/stretchr/testify v1.8.4

## [0.2.4] - 2023-05-11

### Changed in 0.2.4

- Added "origin" to observer messages.
- Updated dependencies
  - github.com/go-sql-driver/mysql v1.7.1
  - github.com/lib/pq v1.10.9
  - github.com/senzing-garage/go-logging v1.2.3
  - github.com/senzing-garage/go-observing v0.2.2

## [0.2.3] - 2023-04-14

### Changed in 0.2.3

- Added "location" to log messages.

## [0.2.2] - 2023-04-13

### Changed in 0.2.2

- Migrated from `github.com/senzing-garage/go-logging/logger` to `github.com/senzing-garage/go-logging/logging`

## [0.2.1] - 2023-02-25

### Added to 0.2.1

- added `connector.NewConnection(ctx, databaseUrl)

## [0.2.0] - 2023-02-24

### Added to 0.2.0

- Added support for Sqlite, Postgresql, MySQL, and MsSQL.
- All support processing SQL via file/bufio.Scanner (sqlexecutor)
- For Postgresql, added GetCurrentWatermark() (postgres)

## [0.1.0] - 2023-02-22

### Added to 0.1.0

- Initial artifacts

[CommonMark]: https://commonmark.org/
[Keep a Changelog]: https://keepachangelog.com/
[Semantic Versioning]: https://semver.org/
