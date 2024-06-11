//go:build darwin

package dbhelper

var testCasesForOsArch = []testCaseMetadata{
	{
		databaseFilename: "/var/opt/senzing/sqlite/G2C.db",
		databaseURL:      "sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db",
		name:             "sqlite3-001",
		succeeds:         true,
	},
	{
		databaseFilename: "/var/opt/senzing/sqlite/G2C.db",
		databaseURL:      `sqlite3://na:na@hostname/var/opt/senzing/sqlite/G2C.db`,
		name:             "sqlite3-002",
		succeeds:         true,
	},
}
