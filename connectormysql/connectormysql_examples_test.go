//go:build linux

package connectormysql

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector() {
	// For more information, visit https://github.com/Senzing/go-databasing/blob/main/connectormysql/connectormysql_examples_test.go
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
