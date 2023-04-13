# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
[markdownlint](https://dlaa.me/markdownlint/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

-

## [0.2.2] - 2023-04-13

### Changed in 0.2.2

- Migrated from `github.com/senzing/go-logging/logger` to `github.com/senzing/go-logging/logging`

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
