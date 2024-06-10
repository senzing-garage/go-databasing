package sqlexecutor

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing-garage/go-databasing/connector"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/stretchr/testify/require"
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

func TestSqlExecutorImpl_ProcessFileName(test *testing.T) {
	ctx := context.TODO()
	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	err := refreshSqliteDatabase(sqliteDatabaseFilename)
	require.NoError(test, err)
	observer1 := &observer.NullObserver{
		ID:       "Observer 1",
		IsSilent: true,
	}
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	if err != nil {
		test.Error(err)
	}
	testObject := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	err = testObject.RegisterObserver(ctx, observer1)
	require.NoError(test, err)
	testObject.SetObserverOrigin(ctx, "Test")
	err = testObject.ProcessFileName(ctx, sqlFilename)
	require.NoError(test, err)
}

func TestSqlExecutorImpl_ProcessScanner(test *testing.T) {
	ctx := context.TODO()
	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	refreshSqliteDatabase(sqliteDatabaseFilename)
	file, err := os.Open(sqlFilename)
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()
	require.NoError(test, err)
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseURL)
	require.NoError(test, err)
	testObject := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	err = testObject.ProcessScanner(ctx, bufio.NewScanner(file))
	require.NoError(test, err)
}
