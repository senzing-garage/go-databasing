//go:build darwin

package dbhelper

import "net/url"

func extractSqliteDatabaseFilenameForOsArch(parsedDatabaseURL *url.URL) (string, error) {
	return parsedDatabaseURL.Path, nil
}
