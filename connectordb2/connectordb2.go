package connectordb2

import (
	"context"
	"database/sql/driver"

	"github.com/ibmdb/go_ibm_db"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// Connector represents a fixed configuration for the pq driver with a given
// name. Connector satisfies the database/sql/driver Connector interface and
// can be used to create any number of DB Conn's via the database/sql OpenDB
// function.
//
// See https://golang.org/pkg/database/sql/driver/#Connector.
// See https://golang.org/pkg/database/sql/#OpenDB.
type Db2 struct {
	Connection string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Connect returns a connection to the database using the fixed configuration
// of this Connector. Context is not used.
func (connector *Db2) Connect(_ context.Context) (driver.Conn, error) {
	return connector.Driver().Open(connector.Connection)
}

// Driver returns the underlying driver of this Connector.
func (connector *Db2) Driver() driver.Driver {
	return &go_ibm_db.Driver{}
}

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
Wrapper for https://pkg.go.dev/github.com/microsoft/go-mssqldb#NewConnector

Input
  - configuration: See https://github.com/microsoft/go-mssqldb
*/
func NewConnector(ctx context.Context, dsn string) (driver.Connector, error) {
	return &Db2{}, nil
}
