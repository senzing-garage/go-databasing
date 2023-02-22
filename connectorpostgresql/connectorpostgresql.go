package connectorpostgresql

import (
	"context"

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
func NewConnector(ctx context.Context, dsn string) (*pq.Connector, error) {
	return pq.NewConnector(dsn)
}
