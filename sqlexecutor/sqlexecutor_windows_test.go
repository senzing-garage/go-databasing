//go:build windows

package sqlexecutor_test

var (
	sqliteDatabaseFilename = `C:\Temp\sqlite\G2C.db`
	sqliteDatabaseURL      = `sqlite3://na:na@nowhere/C:\Temp\sqlite\G2C.db`
)
