package connectorsqlite

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
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connectorsqlite/connectorsqlite_examples_test.go
	ctx := context.TODO()
	configuration := "/tmp/sqlite/G2C.db"
	databaseConnector, err := NewConnector(ctx, configuration)
	printErr(err)
	_ = databaseConnector // Faux use of databaseConnector
	// Output:
}
