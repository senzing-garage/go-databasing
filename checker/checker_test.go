package checker

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing-garage/go-databasing/connector"
	"github.com/senzing-garage/go-databasing/sqlexecutor"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error = nil
	return err
}

func teardown() error {
	var err error = nil
	return err
}

// ----------------------------------------------------------------------------
// Utility functions
// ----------------------------------------------------------------------------

func refreshSqliteDatabase(databaseFilename string) error {
	err := os.Remove(databaseFilename)
	if err != nil {
		fmt.Printf("When removing %s got error: %v\n", databaseFilename, err)
	}
	file, err := os.Create(databaseFilename)
	if err != nil {
		fmt.Printf("When creating %s got error: %v\n", databaseFilename, err)
	}
	file.Close()
	return nil
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestCheckerImpl_IsSchemaInstalled_True(test *testing.T) {
	ctx := context.TODO()

	// Make a fresh database and create Senzing schema.

	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	refreshSqliteDatabase(sqliteDatabaseFilename)
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseUrl)
	if err != nil {
		test.Error(err)
	}
	sqlExecutor := &sqlexecutor.SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	sqlExecutor.ProcessFileName(ctx, sqlFilename)

	// Perform test.

	testObject := &CheckerImpl{
		DatabaseConnector: databaseConnector,
	}
	isSchemaInstalled, err := testObject.IsSchemaInstalled(ctx)
	if isSchemaInstalled {
		if err != nil {
			assert.FailNow(test, err.Error())
		}
	} else {
		if err == nil {
			assert.FailNow(test, "An error should have been returned")
		}
	}
}

func TestCheckerImpl_IsSchemaInstalled_False(test *testing.T) {
	ctx := context.TODO()

	// Make a fresh database with no Senzing schema.

	refreshSqliteDatabase(sqliteDatabaseFilename)

	// Test.

	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseUrl)
	if err != nil {
		test.Error(err)
	}
	testObject := &CheckerImpl{
		DatabaseConnector: databaseConnector,
	}
	isSchemaInstalled, err := testObject.IsSchemaInstalled(ctx)
	if isSchemaInstalled {
		if err != nil {
			assert.FailNow(test, err.Error())
		} else {
			assert.FailNow(test, "Schema is not installed")
		}
	} else {
		if err == nil {
			assert.FailNow(test, "Error should have been returned")
		}
	}

}
