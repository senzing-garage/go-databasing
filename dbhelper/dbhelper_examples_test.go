//go:build linux

package dbhelper

import "fmt"

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleExtractSqliteDatabaseFilename() {
	// For more information, visit https://github.com/senzing-garage/go-databasing/blob/main/dbhelper/dbhelper_examples_test.go
	databaseUrl := "sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db"
	databaseFilename, err := ExtractSqliteDatabaseFilename(databaseUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(databaseFilename)
	// Output: /var/opt/senzing/sqlite/G2C.db
}
