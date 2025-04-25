//go:build linux

package postgresql_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-databasing/connectorpostgresql"
	"github.com/senzing-garage/go-databasing/postgresql"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleBasicPostgresql_GetCurrentWatermark() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/postgresql/postgresql_examples_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
	configuration := "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable"
	databaseConnector, err := connectorpostgresql.NewConnector(ctx, configuration)
	failOnError(err)

	database := &postgresql.BasicPostgresql{
		DatabaseConnector: databaseConnector,
	}
	oid, age, err := database.GetCurrentWatermark(ctx)
	failOnError(err)

	_ = oid // Faux use of oid
	_ = age // Faux use of age
	// Output:
}

func ExampleNewConnector_null() {
	// Output:
} // Hack to format godoc documentation examples properly

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func failOnError(err error) {
	if err != nil {
		fmt.Println(err) //nolint
	}
}
