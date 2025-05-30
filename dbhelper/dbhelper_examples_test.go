//go:build linux

package dbhelper_test

import (
	"fmt"

	"github.com/senzing-garage/go-databasing/dbhelper"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleExtractSqliteDatabaseFilename() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/dbhelper/dbhelper_examples_test.go
	databaseURL := "sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db"
	databaseFilename, err := dbhelper.ExtractSqliteDatabaseFilename(databaseURL)
	failOnError(err)
	fmt.Println(databaseFilename)
	// Output: /var/opt/senzing/sqlite/G2C.db
}

func ExampleGetMessenger() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/dbhelper/dbhelper_examples_test.go
	componentID := 1
	idMessages := map[int]string{}
	callerSkip := 0
	options := []interface{}{}
	_ = dbhelper.GetMessenger(componentID, idMessages, callerSkip, options...)
	// Output:
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func failOnError(err error) {
	if err != nil {
		fmt.Println(err) //nolint
	}
}
