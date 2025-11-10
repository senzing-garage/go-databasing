package connectorsqlite_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-databasing/connectorsqlite"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connectorsqlite/connectorsqlite_examples_test.go
	ctx := context.TODO()
	configuration := "/tmp/sqlite/G2C.db"
	databaseConnector, err := connectorsqlite.NewConnector(ctx, configuration)
	failOnError(err)

	connection, err := databaseConnector.Connect(ctx)
	failOnError(err)

	_ = connection // Faux use of database connection.
	// Output:
}

func ExampleNewConnector_inmemory() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connectorsqlite/connectorsqlite_examples_test.go
	ctx := context.TODO()
	configuration := "/tmp/sqlite/G2C.db?mode=memory&cache=shared"
	databaseConnector, err := connectorsqlite.NewConnector(ctx, configuration)
	failOnError(err)

	connection, err := databaseConnector.Connect(ctx)
	failOnError(err)

	_ = connection // Faux use of database connection.
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
