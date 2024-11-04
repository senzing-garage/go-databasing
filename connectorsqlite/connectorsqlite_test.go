package connectorsqlite

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestNewConnector(test *testing.T) {
	ctx := context.TODO()
	configuration := "/tmp/sqlite/G2C.db"
	databaseConnector, err := NewConnector(ctx, configuration)
	require.NoError(test, err)
	_, err = databaseConnector.Connect(ctx)
	require.NoError(test, err)
}

func TestNewConnectorInMemory(test *testing.T) {
	ctx := context.TODO()
	configuration := "/tmp/sqlite/G2C.db?mode=memory&cache=shared"
	databaseConnector, err := NewConnector(ctx, configuration)
	require.NoError(test, err)
	_, err = databaseConnector.Connect(ctx)
	require.NoError(test, err)
}
