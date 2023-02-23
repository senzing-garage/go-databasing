package connectormssql

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

func TestNewConnector(test *testing.T) {
	ctx := context.TODO()
	configuration := "user id=sa;password=Passw0rd;database=G2;server=localhost"
	databaseConnector, err := NewConnector(ctx, configuration)
	if err != nil {
		assert.FailNow(test, err.Error(), databaseConnector)
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/connectormssql/connectormssql_test.go
	ctx := context.TODO()
	// See https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn
	configuration := "user id=sa;password=Passw0rd;database=master;server=localhost"
	databaseConnector, err := NewConnector(ctx, configuration)
	if err != nil {
		fmt.Println(err, databaseConnector)
	}
	// Output:
}
