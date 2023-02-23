package sqlexecutor

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing/go-databasing/connectorsqlite"
	"github.com/senzing/go-logging/logger"
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
// Test interface functions
// ----------------------------------------------------------------------------

func TestSqlExecutorImpl_ProcessFileName(test *testing.T) {
	ctx := context.TODO()
	observer1 := &observer.ObserverNull{
		Id:       "Observer 1",
		IsSilent: true,
	}
	databaseConnector := &connectorsqlite.Sqlite{
		Filename: "/tmp/sqlite/G2C.db",
	}
	testObject := &SqlExecutorImpl{
		LogLevel:          logger.LevelTrace,
		DatabaseConnector: databaseConnector,
	}
	testObject.RegisterObserver(ctx, observer1)
	testObject.ProcessFileName(ctx, "/opt/senzing/g2/resources/schema/g2core-schema-sqlite-create.sql")
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleSqlExecutorImpl_ProcessFileName() {
	ctx := context.TODO()
	databaseConnector := &connectorsqlite.Sqlite{
		Filename: "/tmp/sqlite/G2C.db",
	}
	testObject := &SqlExecutorImpl{
		LogLevel:          logger.LevelTrace,
		DatabaseConnector: databaseConnector,
	}
	testObject.ProcessFileName(ctx, "/opt/senzing/g2/resources/schema/g2core-schema-sqlite-create.sql")
	// Output:
}
