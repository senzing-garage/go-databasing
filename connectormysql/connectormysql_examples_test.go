//go:build linux

package connectormysql

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNewConnector() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/connectormysql/connectormysql_examples_test.go
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
	printErr(err)
	_ = databaseConnector // Faux use of databaseConnector
	// Output:
}
