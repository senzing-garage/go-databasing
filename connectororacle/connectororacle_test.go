package connectororacle_test

import (
	"testing"

	"github.com/senzing-garage/go-databasing/connectororacle"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestNewConnector(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	configuration := `user="sys" password="Passw0rd" sysdba=true noTimezoneCheck=true connectString="localhost:1521/FREE"`
	databaseConnector, err := connectororacle.NewConnector(ctx, configuration)
	require.NoError(test, err)
	_, err = databaseConnector.Connect(ctx)
	require.NoError(test, err)
}
