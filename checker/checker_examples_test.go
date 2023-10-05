//go:build linux

package checker

import (
	"context"
	"fmt"

	"github.com/senzing/go-databasing/connector"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleSchemaCheckerImpl_IsSchemaInstalled() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/checker/checker_examples_test.go
	ctx := context.TODO()
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseUrl)
	if err != nil {
		fmt.Println(err)
	}
	testObject := &SchemaCheckerImpl{
		DatabaseConnector: databaseConnector,
	}
	isSchemaInstalled, err := testObject.IsSchemaInstalled(ctx)
	fmt.Printf("isSchemaInstalled: %t", isSchemaInstalled)
	// Output: isSchemaInstalled: false
}
