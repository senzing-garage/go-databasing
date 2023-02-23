package main

import (
	"context"
	"database/sql/driver"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/senzing/go-databasing/connectormssql"
	"github.com/senzing/go-databasing/connectormysql"
	"github.com/senzing/go-databasing/connectorpostgresql"
	"github.com/senzing/go-databasing/connectorsqlite"
	"github.com/senzing/go-databasing/postgresql"
	"github.com/senzing/go-databasing/sqlexecutor"
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-observing/observer"
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
	var err error = nil
	var databaseConnector driver.Connector = nil
	var sqlFilename string = ""
	databaseId := Sqlite

	observer1 := &observer.ObserverNull{
		Id:       "Observer 1",
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

	// Choose among different database connectors.

	switch databaseId {
	case Sqlite:
		databaseConnector, err = connectorsqlite.NewConnector(ctx, "/tmp/sqlite/G2C.db")
		sqlFilename = gitRepositoryDir + "/testdata/sqlite/g2core-schema-sqlite-create.sql"

	case Postgresql:
		// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
		databaseConnector, err = connectorpostgresql.NewConnector(ctx, "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable")
		sqlFilename = gitRepositoryDir + "/testdata/postgresql/g2core-schema-postgresql-create.sql"

	case Mysql:
		// See https://pkg.go.dev/github.com/go-sql-driver/mysql#Config
		configuration := &mysql.Config{
			User:      "root",
			Passwd:    "root",
			Net:       "tcp",
			Addr:      "192.168.1.12",
			Collation: "utf8mb4_general_ci",
			DBName:    "G2",
		}
		databaseConnector, err = connectormysql.NewConnector(ctx, configuration)
		sqlFilename = gitRepositoryDir + "/opt/senzing/g2/resources/schema/g2core-schema-mysql-create.sql"

	case Mssql:
		// See https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn
		databaseConnector, err = connectormssql.NewConnector(ctx, "user id=sa;password=Passw0rd;database=G2;server=localhost")
		sqlFilename = gitRepositoryDir + "/opt/senzing/g2/resources/schema/g2core-schema-mssql-create.sql"

	default:
		err = fmt.Errorf("unknown databaseNumber: %d", databaseId)
	}
	if err != nil {
		fmt.Printf("Could not create a database connector. Error: %v", err)
		os.Exit(1)
	}

	// Process file of SQL.

	testObject := &sqlexecutor.SqlExecutorImpl{
		DatabaseConnector: databaseConnector,
	}
	testObject.RegisterObserver(ctx, observer1)
	testObject.SetLogLevel(ctx, logger.LevelTrace)
	err = testObject.ProcessFileName(ctx, sqlFilename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// PostgreSql only tests.

	if databaseId == Postgresql {
		postgresClient := &postgresql.PostgresqlImpl{
			DatabaseConnector: databaseConnector,
		}
		postgresClient.RegisterObserver(ctx, observer1)
		postgresClient.SetLogLevel(ctx, logger.LevelTrace)
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
