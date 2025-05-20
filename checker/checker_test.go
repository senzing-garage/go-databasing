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

var observerSingleton = &observer.NullObserver{
	ID:       "Observer 1",
	IsSilent: true,
}

const sqliteFilename = "../testdata/sqlite/szcore-schema-sqlite-create.sql"

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBasicChecker_IsSchemaInstalled_True(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	sqliteDatabaseURL := getSqliteDatabaseURL(test)
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	require.NoError(test, err)

	sqlExecutor := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqliteFilename)
	require.NoError(test, err)

	// Perform test.

	testObject := getTestObject(ctx, test, sqliteDatabaseURL)

	isSchemaInstalled, err := testObject.IsSchemaInstalled(ctx)
	if isSchemaInstalled {
		require.NoError(test, err)
	} else {
		require.Error(test, err, "An error should have been returned")
	}

	// Perform RecordCount test.

	_, err = testObject.RecordCount(ctx)
	require.NoError(test, err)
}

func TestBasicChecker_RegisterObserver(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test, getSqliteDatabaseURL(test))
	err := testObject.RegisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

func TestBasicChecker_SetLogLevel(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test, getSqliteDatabaseURL(test))
	err := testObject.SetLogLevel(ctx, "DEBUG")
	require.NoError(test, err)
}

func TestBasicChecker_SetLogLevel_badLevelName(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test, getSqliteDatabaseURL(test))
	err := testObject.SetLogLevel(ctx, "BADLEVELNAME")
	require.Error(test, err)
}

func TestBasicChecker_SetObserverOrigin(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test, getSqliteDatabaseURL(test))
	testObject.SetObserverOrigin(ctx, "Test observer origin")
}

func TestBasicChecker_UnregisterObserver(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test, getSqliteDatabaseURL(test))

	// IMPROVE:  This needs to be removed.
	err := testObject.RegisterObserver(ctx, observerSingleton)
	require.NoError(test, err)

	err = testObject.UnregisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

func TestBasicChecker_IsSchemaInstalled_False(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test, getSqliteDatabaseURL(test))
	_, err := testObject.IsSchemaInstalled(ctx)
	require.Error(test, err, "An error should have been returned")
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getBasicChecker(ctx context.Context, t *testing.T, sqliteDatabaseURL string) *checker.BasicChecker {
	t.Helper()

	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	require.NoError(t, err)

	result := &checker.BasicChecker{
		DatabaseConnector: databaseConnector,
	}
	err = result.RegisterObserver(ctx, observerSingleton)
	require.NoError(t, err)
	err = result.SetLogLevel(ctx, "TRACE")
	require.NoError(t, err)

	return result
}

func getDatabaseURL() string {
	ctx := context.Background()

	sqliteDatabaseFile, err := os.CreateTemp(os.TempDir(), "G2C.*.db")
	if err != nil {
		panic(err)
	}

	sqliteDatabaseURL := buildSqliteDatabaseURL(sqliteDatabaseFile.Name())

	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	if err != nil {
		panic(err)
	}

	sqlExecutor := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}

	err = sqlExecutor.ProcessFileName(ctx, sqliteFilename)
	if err != nil {
		panic(err)
	}

	return buildSqliteDatabaseURL(sqliteDatabaseFile.Name())
}

func getSqliteDatabaseURL(t *testing.T) string {
	t.Helper()
	sqliteDatabaseFile, err := os.CreateTemp(t.TempDir(), "G2C.*.db")
	require.NoError(t, err)

	return buildSqliteDatabaseURL(sqliteDatabaseFile.Name())
}

func getTestObject(ctx context.Context, t *testing.T, sqliteDatabaseURL string) checker.Checker {
	t.Helper()

	return getBasicChecker(ctx, t, sqliteDatabaseURL)
}

func failOnError(message string, err error) {
	if err != nil {
		fmt.Printf("%s error: %s", message, err.Error()) //nolint
	}
}
