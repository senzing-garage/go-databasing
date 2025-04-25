//go:build windows

package dbhelper_test

var testCasesForOsArch = []testCaseMetadata{
	{
		databaseFilename: "C:\\Temp\\sqlite\\G2C.db",
		databaseURL:      "sqlite3://na:na@nowhere/C:\\Temp\\sqlite\\G2C.db",
		name:             "sqlite3-001",
		succeeds:         true,
	},
	{
		databaseFilename: `C:\Temp\sqlite\G2C.db`,
		databaseURL:      "sqlite3://na:na@nowhere/C:\\Temp\\sqlite\\G2C.db",
		name:             "sqlite3-002",
		succeeds:         true,
	},
	{
		databaseFilename: "C:\\Temp\\sqlite\\G2C.db",
		databaseURL:      `sqlite3://na:na@nowhere/C:\Temp\sqlite\G2C.db`,
		name:             "sqlite3-003",
		succeeds:         true,
	},
	{
		databaseFilename: `C:\Temp\sqlite\G2C.db`,
		databaseURL:      `sqlite3://na:na@nowhere/C:\Temp\sqlite\G2C.db`,
		name:             "sqlite3-004",
		succeeds:         true,
	},
}
