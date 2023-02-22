package connectormssql

import (
	"context"
	"database/sql/driver"

	mssql "github.com/microsoft/go-mssqldb"
)

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
Wrapper for https://pkg.go.dev/github.com/microsoft/go-mssqldb#NewConnector

Input
  - configuration: See https://github.com/microsoft/go-mssqldb
*/
func NewConnector(ctx context.Context, dsn string) (driver.Connector, error) {
	return mssql.NewConnector(dsn)
}
