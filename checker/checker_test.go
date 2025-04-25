package checker_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing-garage/go-databasing/checker"
	"github.com/senzing-garage/go-databasing/connector"
	"github.com/senzing-garage/go-databasing/sqlexecutor"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/stretchr/testify/require"
)

var (
	observerSingleton = &observer.NullObserver{
		ID:       "Observer 1",
		IsSilent: true,
	}
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBasicChecker_IsSchemaInstalled_True(test *testing.T) {
	ctx := context.TODO()

	// Make a fresh database and create Senzing schema.

	sqlFilename := "../testdata/sqlite/szcore-schema-sqlite-create.sql"
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

	testObject := getTestObject(ctx, test)

	isSchemaInstalled, err := testObject.IsSchemaInstalled(ctx)
	if isSchemaInstalled {
		require.NoError(test, err)
	} else {
		require.Error(test, err, "An error should have been returned")
	}
}

func TestBasicChecker_RecordCount(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObject(ctx, test)
	_, err := testObject.RecordCount(ctx)
	require.NoError(test, err)
}

func TestBasicChecker_RegisterObserver(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObject(ctx, test)
	err := testObject.RegisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

func TestBasicChecker_SetLogLevel(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObject(ctx, test)
	err := testObject.SetLogLevel(ctx, "DEBUG")
	require.NoError(test, err)
}

func TestBasicChecker_SetLogLevel_badLevelName(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObject(ctx, test)
	err := testObject.SetLogLevel(ctx, "BADLEVELNAME")
	require.Error(test, err)
}

func TestBasicChecker_SetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObject(ctx, test)
	testObject.SetObserverOrigin(ctx, "Test observer origin")
}

func TestBasicChecker_UnregisterObserver(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObject(ctx, test)

	// IMPROVE:  This needs to be removed.
	err := testObject.RegisterObserver(ctx, observerSingleton)
	require.NoError(test, err)

	err = testObject.UnregisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

func TestBasicChecker_IsSchemaInstalled_False(test *testing.T) {
	ctx := context.TODO()
	err := refreshSqliteDatabase(sqliteDatabaseFilename)
	require.NoError(test, err)
	testObject := getTestObject(ctx, test)
	_, err = testObject.IsSchemaInstalled(ctx)
	require.Error(test, err, "An error should have been returned")
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getBasicChecker(ctx context.Context, test *testing.T) *checker.BasicChecker {
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	require.NoError(test, err)

	result := &checker.BasicChecker{
		DatabaseConnector: databaseConnector,
	}
	err = result.RegisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
	err = result.SetLogLevel(ctx, "TRACE")
	require.NoError(test, err)

	return result
}

func getTestObject(ctx context.Context, test *testing.T) checker.Checker {
	return getBasicChecker(ctx, test)
}

func outputf(format string, message ...any) {
	fmt.Printf(format, message...) //nolint
}

func refreshSqliteDatabase(databaseFilename string) error {
	err := os.Remove(databaseFilename)
	if err != nil {
		outputf("When removing %s got error: %v\n", databaseFilename, err)
	}

	file, err := os.Create(databaseFilename)
	if err != nil {
		outputf("When creating %s got error: %v\n", databaseFilename, err)
	}

	file.Close()

	return nil
}
