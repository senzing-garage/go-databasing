//go:build windows

package checker_test

func buildSqliteDatabaseURL(filename string) string {
	return "sqlite3://na:na@nowhere/" + filename
}
