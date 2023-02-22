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
// Variables
// ----------------------------------------------------------------------------

// Values updated via "go install -ldflags" parameters.

var (
	programName    string = "unknown"
	buildVersion   string = "0.0.0"
	buildIteration string = "0"
)

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func exampleFunction(ctx context.Context, name string, version string, iteration string) error {
	fmt.Printf("exampleFunction: %s  %s-%s\n", programName, buildVersion, buildIteration)
	return nil
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	ctx := context.TODO()

	databaseConnector := &connectorsqlite.Sqlite{
		Filename: "/tmp/sqlite/G2C.db",
	}
	testObject := &sqlexecutor.SqlExecutorImpl{
		LogLevel:          logger.LevelTrace,
		DatabaseConnector: databaseConnector,
	}
	testObject.ProcessFileName(ctx, "/opt/senzing/g2/resources/schema/g2core-schema-sqlite-create.sql")
}
