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
Function NewConnector is a wrapper for [Microsoft's MSSQL connector].

Input
  - configuration: See [microsoft/go-mssqldb].

[Microsoft's MSSQL connector]: https://pkg.go.dev/github.com/microsoft/go-mssqldb#NewConnector
[microsoft/go-mssqldb]: https://github.com/microsoft/go-mssqldb
*/
func NewConnector(ctx context.Context, dsn string) (driver.Connector, error) {
	_ = ctx
	return mssql.NewConnector(dsn)
}
