package connectorsqlite

import (
	"context"
	"fmt"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connectorsqlite/connectorsqlite_examples_test.go
	ctx := context.TODO()
	configuration := "/tmp/sqlite/G2C.db"
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
