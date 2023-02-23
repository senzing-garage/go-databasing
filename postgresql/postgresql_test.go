package postgresql

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing/go-databasing/connectorpostgresql"
	"github.com/senzing/go-observing/observer"
	"github.com/stretchr/testify/assert"
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

func TestPostgresqlImpl_GetCurrentWatermark(test *testing.T) {
	ctx := context.TODO()
	observer1 := &observer.ObserverNull{
		Id:       "Observer 1",
		IsSilent: true,
	}
	configuration := "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable"
	databaseConnector, err := connectorpostgresql.NewConnector(ctx, configuration)
	if err != nil {
		assert.FailNow(test, err.Error(), databaseConnector)
	}
	testObject := &PostgresqlImpl{
		DatabaseConnector: databaseConnector,
	}
	err = testObject.RegisterObserver(ctx, observer1)
	if err != nil {
		assert.FailNow(test, err.Error())
	}
	oid, age, err := testObject.GetCurrentWatermark(ctx)
	if err != nil {
		assert.FailNow(test, err.Error(), oid, age)
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExamplePostgresqlImpl_GetCurrentWatermark() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/postgresql/postgresql_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
	configuration := "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable"
	databaseConnector, err := connectorpostgresql.NewConnector(ctx, configuration)
	if err != nil {
		fmt.Println(err)
	}
	database := &PostgresqlImpl{
		DatabaseConnector: databaseConnector,
	}
	oid, age, err := database.GetCurrentWatermark(ctx)
	if err != nil {
		fmt.Println(err, oid, age)
	}
	// Output:
}
