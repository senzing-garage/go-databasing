package connectormssql

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
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connectormssql/connectormssql_examples_test.go
	ctx := context.TODO()
	// See https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn
	configuration := "user id=sa;password=Passw0rd;database=master;server=localhost"
	databaseConnector, err := NewConnector(ctx, configuration)
	printErr(err)
	_ = databaseConnector // Faux use of databaseConnector
	// Output:
}
