package postgresql

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing-garage/go-databasing/connectorpostgresql"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestPostgresqlImpl_GetCurrentWatermark(test *testing.T) {
	ctx := context.TODO()
	observer1 := &observer.NullObserver{
		ID:       "Observer 1",
		IsSilent: true,
	}
	configuration := "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable"
	databaseConnector, err := connectorpostgresql.NewConnector(ctx, configuration)
	require.NoError(test, err)
	testObject := &BasicPostgresql{
		DatabaseConnector: databaseConnector,
	}
	err = testObject.RegisterObserver(ctx, observer1)
	require.NoError(test, err)
	testObject.SetObserverOrigin(ctx, "Test")
	_, _, err = testObject.GetCurrentWatermark(ctx)
	require.NoError(test, err)
}

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
	var err error
	return err
}

func teardown() error {
	var err error
	return err
}
