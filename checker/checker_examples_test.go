//go:build linux

package checker

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-databasing/connector"
)

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleBasicChecker_IsSchemaInstalled() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/checker/checker_examples_test.go
	ctx := context.TODO()
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	printErr(err)
	testObject := &BasicChecker{
		DatabaseConnector: databaseConnector,
	}
	isSchemaInstalled, _ := testObject.IsSchemaInstalled(ctx)
	fmt.Printf("isSchemaInstalled: %t", isSchemaInstalled)
	// Output: isSchemaInstalled: false
}
