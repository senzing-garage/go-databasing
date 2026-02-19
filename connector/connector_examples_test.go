//go:build linux

package connector_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-databasing/connector"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector_sqlite() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connector/connector_examples_test.go
	ctx := context.TODO()
	databaseURL := "sqlite3://na:na@$/tmp/sqlite/G2C.db"
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	failOnError(err)

	_ = databaseConnector // Faux use of databaseConnector
	// Output:
}

func ExampleNewConnector_postgresql() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connector/connector_examples_test.go
	ctx := context.TODO()
	databaseURL := "postgresql://postgres:postgres@localhost/G2/?sslmode=disable" //nolint:gosec
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	failOnError(err)

	_ = databaseConnector // Faux use of databaseConnector
	// Output:
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func failOnError(err error) {
	if err != nil {
		fmt.Println(err) //nolint
	}
}
