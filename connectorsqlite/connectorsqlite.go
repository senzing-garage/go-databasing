package connectorsqlite

import (
	"context"
	"database/sql/driver"

	sqlite "github.com/mattn/go-sqlite3"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

/*
Type Sqlite struct implements [driver.Connnector] interface.
This allows Sqlite to be used with [OpenDB].

[driver.Connnector]: https://golang.org/pkg/database/sql/driver/#Connector
[OpenDB]: https://golang.org/pkg/database/sql/#OpenDB
*/
type Sqlite struct {
	Filename string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
Method Connect implements [driver.Connector]'s Connect method.
Context is not used.

[driver.Connector]: https://golang.org/pkg/database/sql/driver/#Connector
*/
func (connector *Sqlite) Connect(_ context.Context) (driver.Conn, error) {
	return connector.Driver().Open(connector.Filename)
}

/*
Method Driver implements [driver.Connector]'s Driver method.

[driver.Connector]: https://golang.org/pkg/database/sql/driver/#Connector
*/
func (connector *Sqlite) Driver() driver.Driver {
	return &sqlite.SQLiteDriver{} //nolint
}

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
Function NewConnector is a wrapper for [github.com/mattn/go-sqlite3].

Input
  - filename: See [github.com/mattn/go-sqlite3].

Output
  - [driver.Connector] configured for SQLite.

[github.com/mattn/go-sqlite3]: https://github.com/mattn/go-sqlite3
[driver.Connector]: https://golang.org/pkg/database/sql/driver/#Connector
*/
func NewConnector(ctx context.Context, filename string) (driver.Connector, error) {
	_ = ctx

	return &Sqlite{
		Filename: filename,
	}, nil
}
