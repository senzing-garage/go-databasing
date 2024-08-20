//go:build linux

package sqlexecutor

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/senzing-garage/go-databasing/connector"
)

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleBasicSQLExecutor_ProcessFileName_mysql() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/go-sql-driver/mysql#Config
	databaseURL := "mysql://root:root@localhost:3306/G2" // #nosec G101
	sqlFilename := "../testdata/mysql/szcore-schema-mysql-create.sql"
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	printErr(err)
	sqlExecutor := &BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	printErr(err)
	// Output:
}

func ExampleBasicSQLExecutor_ProcessFileName_postgresql() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
	databaseURL := "postgresql://postgres:postgres@localhost:5432/G2/?sslmode=disable"
	sqlFilename := "../testdata/postgresql/szcore-schema-postgresql-create.sql"
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	printErr(err)
	sqlExecutor := &BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	printErr(err)
	// Output:
}

func ExampleBasicSQLExecutor_ProcessFileName_mssql() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	// See https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn
	databaseURL := "mssql://sa:Passw0rd@localhost:1433/master"
	sqlFilename := "../testdata/mssql/szcore-schema-mssql-create.sql"
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	printErr(err)
	sqlExecutor := &BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	printErr(err)
	// Output:
}

func ExampleBasicSQLExecutor_ProcessFileName_sqlite() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	databaseFilename := "/tmp/sqlite/G2C.db"
	databaseURL := fmt.Sprintf("sqlite3://na:na@%s", databaseFilename)
	sqlFilename := "../testdata/sqlite/szcore-schema-sqlite-create.sql"
	err := refreshSqliteDatabase(databaseFilename) // Only needed for repeatable test cases.
	printErr(err)
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	printErr(err)
	sqlExecutor := &BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	printErr(err)
	// Output:
}

func ExampleBasicSQLExecutor_ProcessScanner_sqlite() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	databaseFilename := "/tmp/sqlite/G2C.db"
	databaseURL := fmt.Sprintf("sqlite3://na:na@%s", databaseFilename)
	sqlFilename := "../testdata/sqlite/szcore-schema-sqlite-create.sql"
	err := refreshSqliteDatabase(databaseFilename) // Only needed for repeatable test cases.
	printErr(err)
	file, err := os.Open(sqlFilename)
	printErr(err)
	defer file.Close()
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	printErr(err)
	sqlExecutor := &BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessScanner(ctx, bufio.NewScanner(file))
	printErr(err)
	// Output:
}
