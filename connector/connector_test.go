package connector

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestNewConnector_badURLProtocol(test *testing.T) {
	ctx := context.TODO()
	databaseURL := "badProtocol://username:password@hostname:3306/database?schema=schemaname"
	_, err := NewConnector(ctx, databaseURL)
	require.Error(test, err)
}

func TestNewConnector_mysql1(test *testing.T) {
	ctx := context.TODO()
	databaseURL := "mysql://username:password@hostname:3306/database?schema=schemaname"
	_, err := NewConnector(ctx, databaseURL)
	require.NoError(test, err)
}

func TestNewConnector_mysql2(test *testing.T) {
	ctx := context.TODO()
	databaseURL := "mysql://username:password@hostname:3306/?schema=schemaname"
	_, err := NewConnector(ctx, databaseURL)
	require.NoError(test, err)
}

func TestNewConnector_mssql1(test *testing.T) {
	ctx := context.TODO()
	databaseURL := "mssql://username:password@hostname:3306/database?schema=schemaname"
	_, err := NewConnector(ctx, databaseURL)
	require.NoError(test, err)
}

func TestNewConnector_oracle1(test *testing.T) {
	ctx := context.TODO()
	databaseURL := "oracle://sys:Passw0rd@localhost:1521/FREE/?sysdba=true&noTimezoneCheck=true"
	_, err := NewConnector(ctx, databaseURL)
	require.NoError(test, err)
}

func TestNewConnector_postgresql1(test *testing.T) {
	ctx := context.TODO()
	databaseURL := "postgresql://username:password@hostname:5432:database/?schema=schemaname"
	_, err := NewConnector(ctx, databaseURL)
	require.NoError(test, err)
}

func TestNewConnector_postgresql2(test *testing.T) {
	ctx := context.TODO()
	databaseURL := "postgresql://username:password@hostname:5432/database/?schema=schemaname"
	_, err := NewConnector(ctx, databaseURL)
	require.NoError(test, err)
}

func TestNewConnector_sqlite(test *testing.T) {
	ctx := context.TODO()
	databaseURL := "sqlite3://na:na@/tmp/sqlite/G2C.db"
	_, err := NewConnector(ctx, databaseURL)
	require.NoError(test, err)
}
