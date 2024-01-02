//go:build linux

package checker

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-databasing/connector"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleCheckerImpl_IsSchemaInstalled() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/checker/checker_examples_test.go
	ctx := context.TODO()
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseUrl)
	if err != nil {
		fmt.Println(err)
	}
	testObject := &CheckerImpl{
		DatabaseConnector: databaseConnector,
	}
	isSchemaInstalled, err := testObject.IsSchemaInstalled(ctx)
	fmt.Printf("isSchemaInstalled: %t", isSchemaInstalled)
	// Output: isSchemaInstalled: false
}
