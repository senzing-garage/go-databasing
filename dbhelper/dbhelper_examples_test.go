//go:build linux

package dbhelper

import "fmt"

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleExtractSqliteDatabaseFilename() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/dbhelper/dbhelper_examples_test.go
	databaseURL := "sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db"
	databaseFilename, err := ExtractSqliteDatabaseFilename(databaseURL)
	printErr(err)
	fmt.Println(databaseFilename)
	// Output: /var/opt/senzing/sqlite/G2C.db
}
