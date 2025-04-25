//go:build linux

package checker_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-databasing/checker"
	"github.com/senzing-garage/go-databasing/connector"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleBasicChecker_IsSchemaInstalled() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/checker/checker_examples_test.go
	ctx := context.TODO()
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	failOnError(err)

	myChecker := &checker.BasicChecker{
		DatabaseConnector: databaseConnector,
	}
	isSchemaInstalled, _ := myChecker.IsSchemaInstalled(ctx)
	fmt.Printf("isSchemaInstalled: %t", isSchemaInstalled)
	// Output: isSchemaInstalled: false
}

func ExampleBasicChecker_RecordCount() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/checker/checker_examples_test.go
	ctx := context.TODO()
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	failOnError(err)

	myChecker := &checker.BasicChecker{
		DatabaseConnector: databaseConnector,
	}
	recordCount, err := myChecker.RecordCount(ctx)
	failOnError(err)

	_ = recordCount // Faux use of recordCount
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func failOnError(err error) {
	if err != nil {
		fmt.Println(err) //nolint
	}
}
