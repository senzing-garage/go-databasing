# go-databasing

If you are beginning your journey with [Senzing],
please start with [Senzing Quick Start guides].

You are in the [Senzing Garage] where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: go-databasing is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing `go-databasing` packages provide access to the Senzing database.

[![Go Reference Badge]][Package reference]
[![Go Report Card Badge]][Go Report Card]
[![License Badge]][License]
[![go-test-linux.yaml Badge]][go-test-linux.yaml]
[![go-test-darwin.yaml Badge]][go-test-darwin.yaml]
[![go-test-windows.yaml Badge]][go-test-windows.yaml]

[![golangci-lint.yaml Badge]][golangci-lint.yaml]

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

See [main.go] for an example of use.

### Database URLs

## References

1. [API documentation]
1. [Development]
1. [Errors]
1. [Examples]
1. [Package reference]

[API documentation]: https://pkg.go.dev/github.com/senzing-garage/go-databasing
[Development]: docs/development.md
[Errors]: docs/errors.md
[Examples]: docs/examples.md
[Go Reference Badge]: https://pkg.go.dev/badge/github.com/senzing-garage/go-databasing.svg
[Go Report Card Badge]: https://goreportcard.com/badge/github.com/senzing-garage/go-databasing
[Go Report Card]: https://goreportcard.com/report/github.com/senzing-garage/go-databasing
[go-test-darwin.yaml Badge]: https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-darwin.yaml/badge.svg
[go-test-darwin.yaml]: https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-darwin.yaml
[go-test-linux.yaml Badge]: https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-linux.yaml/badge.svg
[go-test-linux.yaml]: https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-linux.yaml
[go-test-windows.yaml Badge]: https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-windows.yaml/badge.svg
[go-test-windows.yaml]: https://github.com/senzing-garage/go-databasing/actions/workflows/go-test-windows.yaml
[golangci-lint.yaml Badge]: https://github.com/senzing-garage/go-databasing/actions/workflows/golangci-lint.yaml/badge.svg
[golangci-lint.yaml]: https://github.com/senzing-garage/go-databasing/actions/workflows/golangci-lint.yaml
[License Badge]: https://img.shields.io/badge/License-Apache2-brightgreen.svg
[License]: https://github.com/senzing-garage/go-databasing/blob/main/LICENSE
[main.go]: main.go
[Package reference]: https://pkg.go.dev/github.com/senzing-garage/go-databasing
[Senzing Garage]: https://github.com/senzing-garage
[Senzing Quick Start guides]: https://docs.senzing.com/quickstart/
[Senzing]: https://senzing.com/
