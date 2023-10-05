//go:build linux

package connector

import (
	"context"
	"fmt"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector_sqlite() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/connectorpostgresql/connectorpostgresql_examples_test.go
	ctx := context.TODO()
	databaseUrl := "sqlite3://na:na@$/tmp/sqlite/G2C.db"
	databaseConnector, err := NewConnector(ctx, databaseUrl)
	if err != nil {
		fmt.Println(err, databaseConnector)
	}
	// Output:
}
