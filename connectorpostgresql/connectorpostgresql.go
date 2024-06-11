package connectorpostgresql

import (
	"context"
	"database/sql/driver"

	"github.com/lib/pq"
)

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
Wrapper for https://pkg.go.dev/github.com/lib/pq#NewConnector

Input
  - dsn: See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
*/
func NewConnector(ctx context.Context, dsn string) (driver.Connector, error) {
	_ = ctx
	return pq.NewConnector(dsn)
}
