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
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connector/connector_examples_test.go
	ctx := context.TODO()
	databaseURL := "sqlite3://na:na@$/tmp/sqlite/G2C.db"
	databaseConnector, err := NewConnector(ctx, databaseURL)
	failOnError(err)
	_ = databaseConnector // Faux use of databaseConnector
	// Output:
}

func ExampleNewConnector_postgresql() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connector/connector_examples_test.go
	ctx := context.TODO()
	databaseURL := "postgresql://postgres:postgres@localhost/G2/?sslmode=disable"
	databaseConnector, err := NewConnector(ctx, databaseURL)
	failOnError(err)
	_ = databaseConnector // Faux use of databaseConnector
	// Output:
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func failOnError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
