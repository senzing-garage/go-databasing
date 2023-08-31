//go:build darwin

package dbhelper

import "net/url"

func extractSqliteDatabaseFilenameForOsArch(parsedDatabaseUrl *url.URL) (string, error) {
	return parsedDatabaseUrl.Path, nil
}
