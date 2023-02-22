package main

import (
	"context"
	"database/sql/driver"
	"fmt"
	"os"
	"time"

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

const MessageIdTemplate = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	ctx := context.TODO()
	var err error = nil
	var databaseConnector driver.Connector = nil
	var sqlFilename string = ""

	observer1 := &observer.ObserverNull{
		Id: "Observer 1",
	}

	// Choose among different database connectors.

	databaseNumber := 2
	switch databaseNumber {
	case 1:
		sqlFilename = "/opt/senzing/g2/resources/schema/g2core-schema-sqlite-create.sql"
		databaseConnector, err = connectorsqlite.NewConnector(ctx, "/tmp/sqlite/G2C.db")
	case 2:
		sqlFilename = "/opt/senzing/g2/resources/schema/g2core-schema-postgresql-create.sql"
		databaseConnector, err = connectorpostgresql.NewConnector(ctx, "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable")
	default:
		err = fmt.Errorf("unknown databaseNumber: %d", databaseNumber)
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

	if databaseNumber == 2 {
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
