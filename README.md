# go-databasing

## :warning: WARNING: go-databasing is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing `go-databasing` packages provide access to the Senzing database.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-databasing.svg)](https://pkg.go.dev/github.com/senzing/go-databasing)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/go-databasing)](https://goreportcard.com/report/github.com/senzing/go-databasing)
[![go-test.yaml](https://github.com/Senzing/go-databasing/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/go-databasing/actions/workflows/go-test.yaml)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/Senzing/go-databasing/blob/main/LICENSE)

## Overview

The `go-database` packages support direct access to the Senzing database on the following database engines:

1. Postgresql
1. Sqlite
1. MySQL
1. MsSQL

Specific uses:

1. **connector:**  Used to transform a database URL into a `database/sql/driver.Connector`
1. **postgresql:**  PostgreSQL-specific calls.  Example: Finding current high-water transaction id.
1. **sqlexecutor:** Used to read a file of SQL and send to database.

## Use

(TODO:)

## References

- [API documentation](https://pkg.go.dev/github.com/senzing/go-databasing)
- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
