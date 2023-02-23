package connectormysql

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
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
	configuration := &mysql.Config{
		User:      "root",
		Passwd:    "root",
		Net:       "tcp",
		Addr:      "localhost",
		Collation: "utf8mb4_general_ci",
		DBName:    "G2",
	}
	databaseConnector, err := NewConnector(ctx, configuration)
	if err != nil {
		assert.FailNow(test, err.Error(), databaseConnector)
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/connectormysql/connectormysql_test.go
	ctx := context.TODO()
	// See https://pkg.go.dev/github.com/go-sql-driver/mysql#Config
	configuration := &mysql.Config{
		User:      "root",
		Passwd:    "root",
		Net:       "tcp",
		Addr:      "localhost",
		Collation: "utf8mb4_general_ci",
		DBName:    "G2",
	}
	databaseConnector, err := NewConnector(ctx, configuration)
	if err != nil {
		fmt.Println(err, databaseConnector)
	}
	// Output:
}
