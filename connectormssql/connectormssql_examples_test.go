package connectormssql_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-databasing/connectormssql"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connectormssql/connectormssql_examples_test.go
	ctx := context.TODO()
	// See https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn
	configuration := "user id=sa;password=Passw0rd;database=master;server=localhost"
	databaseConnector, err := connectormssql.NewConnector(ctx, configuration)
	failOnError(err)

	_ = databaseConnector // Faux use of databaseConnector
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
