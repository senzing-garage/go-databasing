package connectormysql

import (
	"context"
	"database/sql/driver"

	"github.com/go-sql-driver/mysql"
)

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
Wrapper for https://pkg.go.dev/github.com/go-sql-driver/mysql#NewConnector

Input
  - configuration: See https://github.com/go-sql-driver/mysql#dsn-data-source-name
*/
func NewConnector(ctx context.Context, configuration *mysql.Config) (driver.Connector, error) {
	_ = ctx
	return mysql.NewConnector(configuration)
}
