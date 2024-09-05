# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], [markdownlint],
and this project adheres to [Semantic Versioning].

## [Unreleased]

-

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

[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[markdownlint]: https://dlaa.me/markdownlint/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html
