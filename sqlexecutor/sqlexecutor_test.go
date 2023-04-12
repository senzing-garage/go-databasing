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
	databaseFilename := "/tmp/sqlite/G2C.db"
	databaseURL := fmt.Sprintf("sqlite3://na:na@%s", databaseFilename)
	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	refreshSqliteDatabase(databaseFilename)
	observer1 := &observer.ObserverNull{
		Id:       "Observer 1",
		IsSilent: true,
	}
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	if err != nil {
		test.Error(err)
	}
	testObject := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	testObject.RegisterObserver(ctx, observer1)
	testObject.ProcessFileName(ctx, sqlFilename)
}

func TestSqlExecutorImpl_ProcessScanner(test *testing.T) {
	ctx := context.TODO()
	databaseFilename := "/tmp/sqlite/G2C.db"
	databaseURL := fmt.Sprintf("sqlite3://na:na@%s", databaseFilename)
	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	refreshSqliteDatabase(databaseFilename)
	file, err := os.Open(sqlFilename)
	if err != nil {
		test.Error(err)
	}
	defer file.Close()
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	if err != nil {
		test.Error(err)
	}
	testObject := &SqlExecutorImpl{
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
	databaseFilename := "/tmp/sqlite/G2C.db"
	databaseUrl := fmt.Sprintf("sqlite3://na:na@%s", databaseFilename)
	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	refreshSqliteDatabase(databaseFilename) // Only needed for repeatable test cases.
	databaseConnector, err := connector.NewConnector(ctx, databaseUrl)
	if err != nil {
		fmt.Println(err)
	}
	sqlExecutor := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSqlExecutorImpl_ProcessFileName_postgresql() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/sqlexecutor/sqlexecutor_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
	databaseUrl := "postgresql://postgres:postgres@localhost:5432/G2/?sslmode=disable"
	sqlFilename := "../testdata/postgresql/g2core-schema-postgresql-create.sql"
	databaseConnector, err := connector.NewConnector(ctx, databaseUrl)
	if err != nil {
		fmt.Println(err)
	}
	sqlExecutor := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSqlExecutorImpl_ProcessFileName_mssql() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/sqlexecutor/sqlexecutor_test.go
	ctx := context.TODO()
	// See https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn
	databaseUrl := "mssql://sa:Passw0rd@localhost:1433/master"
	sqlFilename := "../testdata/mssql/g2core-schema-mssql-create.sql"
	databaseConnector, err := connector.NewConnector(ctx, databaseUrl)
	if err != nil {
		fmt.Println(err)
	}
	sqlExecutor := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSqlExecutorImpl_ProcessFileName_mysql() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/sqlexecutor/sqlexecutor_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/go-sql-driver/mysql#Config
	databaseUrl := "mysql://root:root@localhost:3306/G2"
	sqlFilename := "../testdata/mysql/g2core-schema-mysql-create.sql"
	databaseConnector, err := connector.NewConnector(ctx, databaseUrl)
	if err != nil {
		fmt.Println(err)
	}
	sqlExecutor := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSqlExecutorImpl_ProcessScanner_sqlite() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/sqlexecutor/sqlexecutor_test.go
	ctx := context.TODO()
	databaseFilename := "/tmp/sqlite/G2C.db"
	databaseUrl := fmt.Sprintf("sqlite3://na:na@%s", databaseFilename)
	sqlFilename := "../testdata/sqlite/g2core-schema-sqlite-create.sql"
	refreshSqliteDatabase(databaseFilename) // Only needed for repeatable test cases.
	file, err := os.Open(sqlFilename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	databaseConnector, err := connector.NewConnector(ctx, databaseUrl)
	if err != nil {
		fmt.Println(err)
	}
	sqlExecutor := &SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessScanner(ctx, bufio.NewScanner(file))
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}
