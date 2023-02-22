package connectormssql

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
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

func TestPostgresql_Connect(test *testing.T) {
	ctx := context.TODO()

	// See https://pkg.go.dev/github.com/go-sql-driver/mysql#Config
	config := &mysql.Config{
		Net:     "tcp",
		Addr:    "1.1.1.1:1234",
		Timeout: 10 * time.Millisecond,
	}

	databaseConnector, err := NewConnector(ctx, config)
	if err != nil {
		test.Fatal(err)

	}
	databaseConnector.Connect(ctx)
}
