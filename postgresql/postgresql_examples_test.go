//go:build linux

package postgresql

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-databasing/connectorpostgresql"
)

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleBasicPostgresql_GetCurrentWatermark() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/postgresql/postgresql_examples_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
	configuration := "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable"
	databaseConnector, err := connectorpostgresql.NewConnector(ctx, configuration)
	printErr(err)
	database := &BasicPostgresql{
		DatabaseConnector: databaseConnector,
	}
	oid, age, err := database.GetCurrentWatermark(ctx)
	printErr(err)
	_ = oid // Faux use of oid
	_ = age // Faux use of age
	// Output:
}
