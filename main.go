package main

import (
	"context"
	"fmt"

	"github.com/senzing/go-databasing/connectorsqlite"
	"github.com/senzing/go-databasing/sqlexecutor"
	"github.com/senzing/go-logging/logger"
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

	databaseConnector, err := connectorsqlite.NewConnector(ctx, "/tmp/sqlite/G2C.db")
	if err != nil {
		fmt.Printf("Could not create a database connector. Error: %v", err)

	}
	testObject := &sqlexecutor.SqlExecutorImpl{
		LogLevel:          logger.LevelTrace,
		DatabaseConnector: databaseConnector,
	}
	testObject.ProcessFileName(ctx, "/opt/senzing/g2/resources/schema/g2core-schema-sqlite-create.sql")
}
