package main

import (
	"context"
	"database/sql/driver"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
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
)

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	ctx := context.TODO()
	var err error = nil
	var databaseConnector driver.Connector = nil
	var sqlFilename string = ""
	databaseId := Mysql

	observer1 := &observer.ObserverNull{
		Id: "Observer 1",
	}

	// Choose among different database connectors.

	switch databaseId {
	case Sqlite:
		databaseConnector, err = connectorsqlite.NewConnector(ctx, "/tmp/sqlite/G2C.db")
		sqlFilename = "/opt/senzing/g2/resources/schema/g2core-schema-sqlite-create.sql"

	case Postgresql:
		// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
		databaseConnector, err = connectorpostgresql.NewConnector(ctx, "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable")
		sqlFilename = "/opt/senzing/g2/resources/schema/g2core-schema-postgresql-create.sql"

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
		sqlFilename = "/opt/senzing/g2/resources/schema/g2core-schema-mysql-create.sql"
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
		}
		fmt.Printf("Postgresql: oid=%s age=%d\n", oid, age)
	}

	// Let Observer finish.

	time.Sleep(2 * time.Second)
}
