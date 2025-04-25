//go:build linux

package sqlexecutor_test

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/senzing-garage/go-databasing/connector"
	"github.com/senzing-garage/go-databasing/sqlexecutor"
)

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
	failOnError(err)

	sqlExecutor := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	failOnError(err)
	// Output:
}

func ExampleBasicSQLExecutor_ProcessFileName_oracle() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/go-sql-driver/mysql#Config
	databaseURL := "oci://sys:Passw0rd@localhost:1521/FREE/?sysdba=true&noTimezoneCheck=true"
	sqlFilename := "../testdata/oracle/szcore-schema-oracle-create.sql"
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	failOnError(err)

	sqlExecutor := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	failOnError(err)
	// Output:
}

func ExampleBasicSQLExecutor_ProcessFileName_postgresql() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
	databaseURL := "postgresql://postgres:postgres@localhost:5432/G2/?sslmode=disable"
	sqlFilename := "../testdata/postgresql/szcore-schema-postgresql-create.sql"
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	failOnError(err)

	sqlExecutor := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	failOnError(err)
	// Output:
}

func ExampleBasicSQLExecutor_ProcessFileName_mssql() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	// See https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn
	databaseURL := "mssql://sa:Passw0rd@localhost:1433/master"
	sqlFilename := "../testdata/mssql/szcore-schema-mssql-create.sql"
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	failOnError(err)

	sqlExecutor := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	failOnError(err)
	// Output:
}

func ExampleBasicSQLExecutor_ProcessFileName_sqlite() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	databaseFilename := "/tmp/sqlite/G2C.db"
	databaseURL := "sqlite3://na:na@" + databaseFilename
	sqlFilename := "../testdata/sqlite/szcore-schema-sqlite-create.sql"

	refreshSqliteDatabase(databaseFilename) // Only needed for repeatable test cases.

	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	failOnError(err)

	sqlExecutor := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	failOnError(err)
	// Output:
}

func ExampleBasicSQLExecutor_ProcessFileName_sqlite_inmemory() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	databaseFilename := "/tmp/sqlite/NotAFile1.db?mode=memory&cache=shared"
	databaseURL := "sqlite3://na:na@" + databaseFilename
	sqlFilename := "../testdata/sqlite/szcore-schema-sqlite-create.sql"
	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	failOnError(err)

	sqlExecutor := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessFileName(ctx, sqlFilename)
	failOnError(err)
	// Output:
}

func ExampleBasicSQLExecutor_ProcessScanner_sqlite() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	databaseFilename := "/tmp/sqlite/G2C.db"
	databaseURL := "sqlite3://na:na@" + databaseFilename
	sqlFilename := "../testdata/sqlite/szcore-schema-sqlite-create.sql"

	refreshSqliteDatabase(databaseFilename) // Only needed for repeatable test cases.

	file, err := os.Open(sqlFilename)
	failOnError(err)

	defer file.Close()

	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	failOnError(err)

	sqlExecutor := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessScanner(ctx, bufio.NewScanner(file))
	failOnError(err)
	// Output:
}

func ExampleBasicSQLExecutor_ProcessScanner_sqlite_inmemory() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/sqlexecutor/sqlexecutor_examples_test.go
	ctx := context.TODO()
	databaseFilename := "/tmp/sqlite/NotAFile2.db?mode=memory&cache=shared"
	databaseURL := "sqlite3://na:na@" + databaseFilename
	sqlFilename := "../testdata/sqlite/szcore-schema-sqlite-create.sql"
	file, err := os.Open(sqlFilename)
	failOnError(err)

	defer file.Close()

	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	failOnError(err)

	sqlExecutor := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = sqlExecutor.ProcessScanner(ctx, bufio.NewScanner(file))
	failOnError(err)
	// Output:
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func failOnError(err error) {
	if err != nil {
		fmt.Println(err) //nolint
	}
}
