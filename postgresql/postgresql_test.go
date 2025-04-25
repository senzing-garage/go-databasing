package postgresql_test

import (
	"context"
	"testing"

	"github.com/senzing-garage/go-databasing/connectorpostgresql"
	"github.com/senzing-garage/go-databasing/postgresql"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/stretchr/testify/require"
)

var (
	observerSingleton = &observer.NullObserver{
		ID:       "Observer 1",
		IsSilent: true,
	}
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBasicPostgresql_GetCurrentWatermark(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test)
	_, _, err := testObject.GetCurrentWatermark(ctx)
	require.NoError(test, err)
}

func TestBasicPostgresql_RegisterObserver(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test)
	err := testObject.RegisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

func TestBasicPostgresql_SetLogLevel(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test)
	err := testObject.SetLogLevel(ctx, "DEBUG")
	require.NoError(test, err)
}

func TestBasicChecker_SetLogLevel_badLevelName(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test)
	err := testObject.SetLogLevel(ctx, "BADLEVELNAME")
	require.Error(test, err)
}

func TestBasicPostgresql_SetObserverOrigin(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test)
	testObject.SetObserverOrigin(ctx, "Test observer origin")
}

func TestBasicPostgresql_UnregisterObserver(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := getTestObject(ctx, test)

	// IMPROVE:  This needs to be removed.
	err := testObject.RegisterObserver(ctx, observerSingleton)
	require.NoError(test, err)

	err = testObject.UnregisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, t *testing.T) postgresql.Postgresql {
	t.Helper()

	return getBasicPostgresql(ctx, t)
}

func getBasicPostgresql(ctx context.Context, t *testing.T) *postgresql.BasicPostgresql {
	t.Helper()

	configuration := "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable"
	databaseConnector, err := connectorpostgresql.NewConnector(ctx, configuration)
	require.NoError(t, err)

	result := &postgresql.BasicPostgresql{
		DatabaseConnector: databaseConnector,
	}
	err = result.RegisterObserver(ctx, observerSingleton)
	require.NoError(t, err)
	err = result.SetLogLevel(ctx, "TRACE")
	require.NoError(t, err)

	return result
}
