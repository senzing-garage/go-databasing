package connectororacle

import (
	"context"
	"database/sql/driver"

	"github.com/godror/godror"
	"github.com/senzing-garage/go-helpers/wraperror"
)

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
Function NewConnector is a wrapper for [Go DRiver for ORacle].

Input
  - dsn: See [Connection strings].

Output
  - [driver.Connector] configured for Oracle.

[Go DRiver for ORacle]: https://github.com/godror/godror
[driver.Connector]: https://pkg.go.dev/database/sql/driver#Connector
[Connection strings]: https://godror.github.io/godror/doc/connection.html
*/
func NewConnector(ctx context.Context, dsn string) (driver.Connector, error) {
	_ = ctx
	params, err := godror.ParseDSN(dsn)

	return godror.NewConnector(params), wraperror.Errorf(err, wraperror.NoMessage)
}
