package sqlexecutor

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing/go-databasing/connector"
	"github.com/senzing/go-observing/observer"
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
	refreshSqliteDatabase(sqliteDatabaseFilename)
	observer1 := &observer.ObserverNull{
		Id:       "Observer 1",
		IsSilent: true,
	}
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseUrl)
	if err != nil {
		test.Error(err)
	}
	testObject := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	testObject.RegisterObserver(ctx, observer1)
	testObject.SetObserverOrigin(ctx, "Test")
	testObject.ProcessFileName(ctx, sqlFilename)
}

func TestSqlExecutorImpl_ProcessScanner(test *testing.T) {
	ctx := context.TODO()
	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	refreshSqliteDatabase(sqliteDatabaseFilename)
	file, err := os.Open(sqlFilename)
	if err != nil {
		test.Error(err)
	}
	defer file.Close()
	databaseConnector, err := connector.NewConnector(ctx, sqliteDatabaseUrl)
	if err != nil {
		test.Error(err)
	}
	testObject := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	testObject.ProcessScanner(ctx, bufio.NewScanner(file))
}
