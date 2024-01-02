# go-databasing

If you are beginning your journey with
[Senzing](https://senzing.com/),
please start with
[Senzing Quick Start guides](https://docs.senzing.com/quickstart/).

You are in the
[Senzing Garage](https://github.com/senzing-garage)
where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: go-databasing is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing `go-databasing` packages provide access to the Senzing database.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing-garage/go-databasing.svg)](https://pkg.go.dev/github.com/senzing-garage/go-databasing)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing-garage/go-databasing)](https://goreportcard.com/report/github.com/senzing-garage/go-databasing)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/senzing-garage/go-databasing/blob/main/LICENSE)

[![gosec.yaml](https://github.com/senzing-garage/go-databasing/actions/workflows/gosec.yaml/badge.svg)](https://github.com/senzing-garage/go-databasing/actions/workflows/gosec.yaml)
[![go-test-linux.yaml](https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-linux.yaml/badge.svg)](https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-linux.yaml)
[![go-test-darwin.yaml](https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-darwin.yaml/badge.svg)](https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-darwin.yaml)
[![go-test-windows.yaml](https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-windows.yaml/badge.svg)](https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-windows.yaml)

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

- [API documentation](https://pkg.go.dev/github.com/senzing-garage/go-databasing)
- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
