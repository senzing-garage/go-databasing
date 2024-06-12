package connectorsqlite

import (
	"context"
	"database/sql/driver"

	sqlite "github.com/mattn/go-sqlite3"
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
type Sqlite struct {
	Filename string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Connect returns a connection to the database using the fixed configuration
// of this Connector. Context is not used.
func (connector *Sqlite) Connect(_ context.Context) (driver.Conn, error) {
	return connector.Driver().Open(connector.Filename)
}

// Driver returns the underlying driver of this Connector.
func (connector *Sqlite) Driver() driver.Driver {
	return &sqlite.SQLiteDriver{}
}

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

func NewConnector(ctx context.Context, filename string) (driver.Connector, error) {
	_ = ctx
	return &Sqlite{
		Filename: filename,
	}, nil
}
