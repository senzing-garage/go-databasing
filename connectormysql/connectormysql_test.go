package connectormysql

import (
	"context"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestNewConnector(test *testing.T) {
	ctx := context.TODO()
	configuration := &mysql.Config{
		User:      "root",
		Passwd:    "root",
		Net:       "tcp",
		Addr:      "localhost",
		Collation: "utf8mb4_general_ci",
		DBName:    "G2",
	}
	databaseConnector, err := NewConnector(ctx, configuration)
	require.NoError(test, err)
	_, err = databaseConnector.Connect(ctx)
	require.NoError(test, err)
}
