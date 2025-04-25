package connectormssql_test

import (
	"context"
	"testing"

	"github.com/senzing-garage/go-databasing/connectormssql"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestNewConnector(test *testing.T) {
	ctx := context.TODO()
	configuration := "user id=sa;password=Passw0rd;database=master;server=localhost"
	databaseConnector, err := connectormssql.NewConnector(ctx, configuration)
	require.NoError(test, err)
	_, err = databaseConnector.Connect(ctx)
	require.NoError(test, err)
}
