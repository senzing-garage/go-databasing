package connectormysql

import (
	"context"
	"database/sql/driver"

	"github.com/go-sql-driver/mysql"
	"github.com/senzing-garage/go-helpers/wraperror"
)

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
Function NewConnector is a wrapper for [go-sql-driver/mysql].

Input
  - configuration: See [DSN (Data Source Name)].

Output
  - [driver.Connector] configured for MySQL.

[DSN (Data Source Name)]: https://github.com/go-sql-driver/mysql#dsn-data-source-name
[go-sql-driver/mysql]: https://pkg.go.dev/github.com/go-sql-driver/mysql#NewConnector
*/
func NewConnector(ctx context.Context, configuration *mysql.Config) (driver.Connector, error) {
	_ = ctx

	result, err := mysql.NewConnector(configuration)

	return result, wraperror.Errorf(err, "connectormysql.NewConnector error: %w", err)
}
