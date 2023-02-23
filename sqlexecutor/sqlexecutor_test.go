package sqlexecutor

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing/go-databasing/connectorpostgresql"
	"github.com/senzing/go-databasing/connectorsqlite"
	"github.com/senzing/go-logging/logger"
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
	databaseFilename := "/tmp/sqlite/G2C.db"
	refreshSqliteDatabase(databaseFilename)
	observer1 := &observer.ObserverNull{
		Id:       "Observer 1",
		IsSilent: true,
	}
	databaseConnector := &connectorsqlite.Sqlite{
		Filename: databaseFilename,
	}
	testObject := &SqlExecutorImpl{
		LogLevel:          logger.LevelInfo,
		DatabaseConnector: databaseConnector,
	}
	testObject.RegisterObserver(ctx, observer1)
	testObject.ProcessFileName(ctx, sqlFilename)
}

func TestSqlExecutorImpl_ProcessScanner(test *testing.T) {
	ctx := context.TODO()
	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	databaseFilename := "/tmp/sqlite/G2C.db"
	refreshSqliteDatabase(databaseFilename)
	file, err := os.Open(sqlFilename)
	if err != nil {
		test.Error(err)
	}
	defer file.Close()
	databaseConnector := &connectorsqlite.Sqlite{
		Filename: databaseFilename,
	}
	testObject := &SqlExecutorImpl{
		LogLevel:          logger.LevelInfo,
		DatabaseConnector: databaseConnector,
	}
	testObject.ProcessScanner(ctx, bufio.NewScanner(file))
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleSqlExecutorImpl_ProcessFileName_sqlite() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/sqlexecutor/sqlexecutor_test.go
	ctx := context.TODO()
	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	databaseFilename := "/tmp/sqlite/G2C.db"
	refreshSqliteDatabase(databaseFilename) // Only needed for repeatable test cases.
	databaseConnector, err := connectorsqlite.NewConnector(ctx, databaseFilename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	sqlExecutor := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	// Output:
}

func ExampleSqlExecutorImpl_ProcessFileName_postgresql() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/sqlexecutor/sqlexecutor_test.go
	ctx := context.TODO()
	sqlFilename := "../testdata/postgresql/g2core-schema-postgresql-create.sql"
	dsn := "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable"
	databaseConnector, err := connectorpostgresql.NewConnector(ctx, dsn)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	sqlExecutor := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	// Output:
}

func ExampleSqlExecutorImpl_ProcessScanner_sqlite() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/sqlexecutor/sqlexecutor_test.go
	ctx := context.TODO()
	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	databaseFilename := "/tmp/sqlite/G2C.db"
	refreshSqliteDatabase(databaseFilename) // Only needed for repeatable test cases.
	file, err := os.Open(sqlFilename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer file.Close()
	databaseConnector, err := connectorsqlite.NewConnector(ctx, databaseFilename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	testObject := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	err = testObject.ProcessScanner(ctx, bufio.NewScanner(file))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	// Output:
}
