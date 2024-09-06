package connectororacle

import (
	"context"
	"fmt"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connectormssql/connectororacle_examples_test.go
	ctx := context.TODO()
	// See https://godror.github.io/godror/doc/connection.html
	// See https://pkg.go.dev/github.com/godror/godror
	// configuration := `user="sysdba" password="Passw0rd" connectString="localhost:1521/FREEPDB1"`
	configuration := `user="sys" sysdba=true password="Passw0rd" connectString="localhost:1521/FREE"`
	databaseConnector, err := NewConnector(ctx, configuration)
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
		fmt.Println(err)
	}
}
