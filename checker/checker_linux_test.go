//go:build linux

package checker_test

func buildSqliteDatabaseURL(filename string) string {
	return "sqlite3://na:na@" + filename
}
