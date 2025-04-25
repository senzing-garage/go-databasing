package connectororacle_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-databasing/connectororacle"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connectormssql/connectororacle_examples_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/godror/godror
	configuration := `user="sys" sysdba=true password="Passw0rd" connectString="localhost:1521/FREE"`
	databaseConnector, err := connectororacle.NewConnector(ctx, configuration)
	failOnError(err)

	_ = databaseConnector // Faux use of databaseConnector
	// Output:
}

func ExampleNewConnector_null() {} // Hack to format godoc documentation examples properly

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func failOnError(err error) {
	if err != nil {
		fmt.Println(err) //nolint
	}
}
