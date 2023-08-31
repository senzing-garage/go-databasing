//go:build windows

package dbhelper

import (
	"net/url"
	"strings"
)

func extractSqliteDatabaseFilenameForOsArch(parsedDatabaseUrl *url.URL) (string, error) {
	result := parsedDatabaseUrl.Path
	if strings.HasPrefix(result, "/") {
		result = result[1:]
	}
	return result, nil
}
