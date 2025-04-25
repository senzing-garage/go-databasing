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
	failOnError("connector.NewConnector", err)

	myChecker := &checker.BasicChecker{
		DatabaseConnector: databaseConnector,
	}
	isSchemaInstalled, _ := myChecker.IsSchemaInstalled(ctx)
	fmt.Printf("isSchemaInstalled: %t", isSchemaInstalled)
	// Output: isSchemaInstalled: true
}

func ExampleBasicChecker_RecordCount() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/checker/checker_examples_test.go
	ctx := context.TODO()

	// err := refreshSqliteDatabase(sqliteDatabaseFilename)
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	failOnError("connector.NewConnector", err)

	myChecker := &checker.BasicChecker{
		DatabaseConnector: databaseConnector,
	}
	recordCount, err := myChecker.RecordCount(ctx)
	failOnError("myChecker.RecordCount", err)

	_ = recordCount // Faux use of recordCount

	// Output:
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func failOnError(message string, err error) {
	if err != nil {
		fmt.Printf("%s error: %s", message, err.Error()) //nolint
	}
}
