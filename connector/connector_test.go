package connector

import (
	"context"
	"fmt"
	"os"
	"testing"

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

func TestNewConnector_sqlite(test *testing.T) {
	ctx := context.TODO()
	databaseUrl := "sqlite3://na:na@/tmp/sqlite/G2C.db"
	databaseConnector, err := NewConnector(ctx, databaseUrl)
	if err != nil {
		assert.FailNow(test, err.Error(), databaseConnector)
	}
}

func TestNewConnector_postgresql1(test *testing.T) {
	ctx := context.TODO()
	databaseUrl := "postgresql://username:password@hostname:5432:database/?schema=schemaname"
	databaseConnector, err := NewConnector(ctx, databaseUrl)
	if err != nil {
		assert.FailNow(test, err.Error(), databaseConnector)
	}
}

func TestNewConnector_postgresql2(test *testing.T) {
	ctx := context.TODO()
	databaseUrl := "postgresql://username:password@hostname:5432/database/?schema=schemaname"
	databaseConnector, err := NewConnector(ctx, databaseUrl)
	if err != nil {
		assert.FailNow(test, err.Error(), databaseConnector)
	}
}

func TestNewConnector_mysql1(test *testing.T) {
	ctx := context.TODO()
	databaseUrl := "mysql://username:password@hostname:3306/database?schema=schemaname"
	databaseConnector, err := NewConnector(ctx, databaseUrl)
	if err != nil {
		assert.FailNow(test, err.Error(), databaseConnector)
	}
}

func TestNewConnector_mysql2(test *testing.T) {
	ctx := context.TODO()
	databaseUrl := "mysql://username:password@hostname:3306/?schema=schemaname"
	databaseConnector, err := NewConnector(ctx, databaseUrl)
	if err != nil {
		assert.FailNow(test, err.Error(), databaseConnector)
	}
}

func TestNewConnector_mssql(test *testing.T) {
	ctx := context.TODO()
	databaseUrl := "mssql://username:password@hostname:3306/database?schema=schemaname"
	databaseConnector, err := NewConnector(ctx, databaseUrl)
	if err != nil {
		assert.FailNow(test, err.Error(), databaseConnector)
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector_sqlite() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/connectorpostgresql/connectorpostgresql_test.go
	ctx := context.TODO()
	databaseUrl := "sqlite3://na:na@$/tmp/sqlite/G2C.db"
	databaseConnector, err := NewConnector(ctx, databaseUrl)
	if err != nil {
		fmt.Println(err, databaseConnector)
	}
	// Output: sqlite3: /tmp/sqlite/G2C.db
}
