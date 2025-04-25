package connectorpostgresql_test

import (
	"testing"

	"github.com/senzing-garage/go-databasing/connectorpostgresql"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestNewConnector(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	configuration := "user=postgres password=postgres dbname=G2 host=localhost sslmode=disable"
	databaseConnector, err := connectorpostgresql.NewConnector(ctx, configuration)
	require.NoError(test, err)
	_, err = databaseConnector.Connect(ctx)
	require.NoError(test, err)
}
