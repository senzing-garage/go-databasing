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
	Sqlite int = iota
	Postgresql
	Mysql
	Mssql
)

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	ctx := context.TODO()
	var sqlFilename string
	var databaseURL string
	databaseID := Sqlite

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
	case Sqlite:
		databaseURL = sqliteDatabaseURL
		sqlFilename = gitRepositoryDir + "/testdata/sqlite/g2core-schema-sqlite-create.sql"
	case Postgresql:
		// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
		databaseURL = "postgresql://postgres:postgres@localhost/G2/?sslmode=disable"
		sqlFilename = gitRepositoryDir + "/testdata/postgresql/g2core-schema-postgresql-create.sql"
	case Mysql:
		// See https://pkg.go.dev/github.com/go-sql-driver/mysql#Config
		databaseURL = "mysql://root:root@localhost/G2" // #nosec G101
		sqlFilename = gitRepositoryDir + "/testdata/mysql/g2core-schema-mysql-create.sql"
	case Mssql:
		// See https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn
		databaseURL = "mysql://sa:Passw0rd@localhost/master"
		sqlFilename = gitRepositoryDir + "/testdata/mssql/g2core-schema-mssql-create.sql"
	default:
		fmt.Printf("Unknown databaseNumber: %d", databaseID)
		os.Exit(1)
	}

	// Create database connector.

	databaseConnector, err := connector.NewConnector(ctx, databaseURL)
	if err != nil {
		fmt.Printf("Could not create a database connector. Error: %v", err)
		os.Exit(1)
	}

	// Process file of SQL.

	testObject := &sqlexecutor.BasicSQLExecutor{
		DatabaseConnector: databaseConnector,
	}
	err = testObject.RegisterObserver(ctx, observer1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	err = testObject.ProcessFileName(ctx, sqlFilename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// PostgreSql only tests.

	if databaseID == Postgresql {
		postgresClient := &postgresql.BasicPostgresql{
			DatabaseConnector: databaseConnector,
		}
		err = postgresClient.RegisterObserver(ctx, observer1)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		err = postgresClient.SetLogLevel(ctx, logging.LevelTraceName)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		oid, age, err := postgresClient.GetCurrentWatermark(ctx)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Postgresql: oid=%s age=%d\n", oid, age)
	}

	// Let Observer finish.

	time.Sleep(2 * time.Second)
}
