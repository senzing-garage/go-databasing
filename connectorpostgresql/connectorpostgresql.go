package connectorpostgresql

import (
	"context"
	"database/sql/driver"

	"github.com/lib/pq"
	"github.com/senzing-garage/go-helpers/wraperror"
)

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
Function NewConnector is a wrapper for [github.com/lib/pq].

Input
  - dsn: See [Connection String Parameters].

Output
  - [driver.Connector] configured for PostgreSQL.

[github.com/lib/pq]: https://pkg.go.dev/github.com/lib/pq#NewConnector
[Connection String Parameters]: https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
*/
func NewConnector(ctx context.Context, dsn string) (driver.Connector, error) {
	_ = ctx

	result, err := pq.NewConnector(dsn)

	return result, wraperror.Errorf(err, "connectorpostgresql.NewConnector error: %w", err)
}
