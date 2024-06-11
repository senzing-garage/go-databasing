package connectorpostgresql

import (
	"context"
	"fmt"
)

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connectorpostgresql/connectorpostgresql_examples_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
	configuration := "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable"
	databaseConnector, err := NewConnector(ctx, configuration)
	printErr(err)
	_ = databaseConnector // Faux use of databaseConnector
	// Output:
}
