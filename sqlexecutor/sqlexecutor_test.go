package sqlexecutor_test

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing-garage/go-databasing/connector"
	"github.com/senzing-garage/go-databasing/sqlexecutor"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/stretchr/testify/require"
)

const (
	sqlFilename = "../testdata/sqlite/szcore-schema-sqlite-create.sql"
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

func TestBasicSQLExecutor_ProcessFileName(test *testing.T) {
	ctx := context.TODO()
	err := refreshSqliteDatabase(sqliteDatabaseFilename)
	require.NoError(test, err)
	testObject := getTestObject(ctx, test)
	err = testObject.ProcessFileName(ctx, sqlFilename)
	require.NoError(test, err)
}

func TestBasicSQLExecutor_ProcessScanner(test *testing.T) {
	ctx := context.TODO()
	err := refreshSqliteDatabase(sqliteDatabaseFilename)
	require.NoError(test, err)
	file, err := os.Open(sqlFilename)

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()
	require.NoError(test, err)
	testObject := getTestObject(ctx, test)
	err = testObject.ProcessScanner(ctx, bufio.NewScanner(file))
	require.NoError(test, err)
}

func TestBasicSQLExecutor_RegisterObserver(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObject(ctx, test)
	err := testObject.RegisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

func TestBasicSQLExecutor_SetLogLevel(test *testing.T) {
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

func TestBasicSQLExecutor_SetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObject(ctx, test)
	testObject.SetObserverOrigin(ctx, "Test observer origin")
}

func TestBasicSQLExecutor_UnregisterObserver(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObject(ctx, test)
	err := testObject.UnregisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getBasicSQLExecutor(ctx context.Context, test *testing.T) *sqlexecutor.BasicSQLExecutor {
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	require.NoError(test, err)

	result := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = result.RegisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
	err = result.SetLogLevel(ctx, "TRACE")
	require.NoError(test, err)

	return result
}

func getTestObject(ctx context.Context, test *testing.T) sqlexecutor.SQLExecutor {
	return getBasicSQLExecutor(ctx, test)
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
