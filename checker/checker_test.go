package checker

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing-garage/go-databasing/connector"
	"github.com/senzing-garage/go-databasing/sqlexecutor"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBasicChecker_IsSchemaInstalled_True(test *testing.T) {
	ctx := context.TODO()

	// Make a fresh database and create Senzing schema.

	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	err := refreshSqliteDatabase(sqliteDatabaseFilename)
	require.NoError(test, err)
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	require.NoError(test, err)
	sqlExecutor := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	require.NoError(test, err)

	// Perform test.

	testObject := &BasicChecker{
		DatabaseConnector: databaseConnector,
	}
	isSchemaInstalled, err := testObject.IsSchemaInstalled(ctx)
	if isSchemaInstalled {
		require.NoError(test, err)
	} else {
		require.Error(test, err, "An error should have been returned")
	}
}

func TestBasicChecker_IsSchemaInstalled_False(test *testing.T) {
	ctx := context.TODO()
	err := refreshSqliteDatabase(sqliteDatabaseFilename)
	require.NoError(test, err)
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	require.NoError(test, err)
	testObject := &BasicChecker{
		DatabaseConnector: databaseConnector,
	}
	_, err = testObject.IsSchemaInstalled(ctx)
	require.Error(test, err, "An error should have been returned")
}

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
	var err error
	return err
}

func teardown() error {
	var err error
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
