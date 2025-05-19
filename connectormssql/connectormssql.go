package connectormssql

import (
	"context"
	"database/sql/driver"

	mssql "github.com/microsoft/go-mssqldb"
	"github.com/senzing-garage/go-helpers/wraperror"
)

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
Function NewConnector is a wrapper for [Microsoft's MS-SQL connector].

Input
  - dsn: See [microsoft/go-mssqldb].

Output
  - [driver.Connector] configured for MS-SQL.

[Microsoft's MS-SQL connector]: https://pkg.go.dev/github.com/microsoft/go-mssqldb#NewConnector
[microsoft/go-mssqldb]: https://github.com/microsoft/go-mssqldb
*/
func NewConnector(ctx context.Context, dsn string) (driver.Connector, error) {
	_ = ctx

	result, err := mssql.NewConnector(dsn)

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}
