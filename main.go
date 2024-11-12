package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/senzing-garage/go-databasing/connector"
	"github.com/senzing-garage/go-databasing/postgresql"
	"github.com/senzing-garage/go-databasing/sqlexecutor"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const (
	Mssql int = iota
	Mysql
	Oracle
	Postgresql
	Sqlite
	SqliteInMemory
)

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	databaseIDs := []int{Oracle, Mssql, Mysql, Sqlite, SqliteInMemory, Postgresql}
	printStatementTemplate := "\n==== %11s ==========================\n\n"
	for _, databaseID := range databaseIDs {
		switch databaseID {
		case Mssql:
			fmt.Printf(printStatementTemplate, "Mssql")
		case Mysql:
			fmt.Printf(printStatementTemplate, "Mysql")
		case Oracle:
			fmt.Printf(printStatementTemplate, "Oracle")
		case Postgresql:
			fmt.Printf(printStatementTemplate, "Postgresql")
		case Sqlite:
			fmt.Printf(printStatementTemplate, "Sqlite")
		case SqliteInMemory:
			fmt.Printf(printStatementTemplate, "SqliteInMemory")
		}
		demonstrateDatabase(databaseID)
	}
}

func demonstrateDatabase(databaseID int) {
	ctx := context.TODO()
	var sqlFilename string
	var databaseURL string

	if databaseID == Postgresql {
		time.Sleep(2 * time.Minute)
	}

	// Create a silent observer.

	observer1 := &observer.NullObserver{
		ID:       "Observer 1",
		IsSilent: true,
	}

	// Get location of test data.

	// gitRepositoryDir, found := os.LookupEnv("GITHUB_WORKSPACE") // For GitHub actions.
	// if !found {
	// 	gitRepositoryDir, found = os.LookupEnv("GIT_REPOSITORY_DIR")
	// }
	// if !found {
	// 	gitRepositoryDir = "."
	// }

	gitRepositoryDir := "."

	// Construct database URL and choose SQL file.

	switch databaseID {
	case Mysql:
		// See https://pkg.go.dev/github.com/go-sql-driver/mysql#Config
		databaseURL = "mysql://root:root@localhost/G2" // #nosec G101
		sqlFilename = gitRepositoryDir + "/testdata/mysql/szcore-schema-mysql-create.sql"
	case Mssql:
		// See https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn
		databaseURL = "mssql://sa:Passw0rd@localhost/master"
		sqlFilename = gitRepositoryDir + "/testdata/mssql/szcore-schema-mssql-create.sql"
	case Oracle:
		// See https://pkg.go.dev/github.com/godror/godror#pkg-overview
		databaseURL = "oracle://sys:Passw0rd@localhost:1521/FREE/?sysdba=true&noTimezoneCheck=true"
		sqlFilename = gitRepositoryDir + "/testdata/oracle/szcore-schema-oracle-create.sql"
	case Postgresql:
		// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
		databaseURL = "postgresql://postgres:postgres@localhost/G2/?sslmode=disable"
		sqlFilename = gitRepositoryDir + "/testdata/postgresql/szcore-schema-postgresql-create.sql"
	case Sqlite:
		databaseURL = sqliteDatabaseURL
		sqlFilename = gitRepositoryDir + "/testdata/sqlite/szcore-schema-sqlite-create.sql"
	case SqliteInMemory:
		databaseURL = sqliteDatabaseURL + "?mode=memory&cache=shared"
		sqlFilename = gitRepositoryDir + "/testdata/sqlite/szcore-schema-sqlite-create.sql"
	default:
		exitOnError(fmt.Errorf("unknown databaseNumber: %d", databaseID))
	}

	// Create database connector.

	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	exitOnError(err)

	// Process file of SQL.

	testObject := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = testObject.RegisterObserver(ctx, observer1)
	exitOnError(err)

	err = testObject.ProcessFileName(ctx, sqlFilename)
	exitOnError(err)

	// PostgreSql only tests.

	if databaseID == Postgresql {
		postgresClient := &postgresql.BasicPostgresql{
			DatabaseConnector: databaseConnector,
		}
		err = postgresClient.RegisterObserver(ctx, observer1)
		exitOnError(err)

		err = postgresClient.SetLogLevel(ctx, logging.LevelTraceName)
		exitOnError(err)

		oid, age, err := postgresClient.GetCurrentWatermark(ctx)
		exitOnError(err)

		fmt.Printf("Postgresql: oid=%s age=%d\n", oid, age)
	}

	// Let Observer finish.

	time.Sleep(2 * time.Second)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
