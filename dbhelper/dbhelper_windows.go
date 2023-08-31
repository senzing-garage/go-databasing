//go:build windows

package dbhelper

import (
	"fmt"
	"net/url"
	"strings"
)

func extractSqliteDatabaseFilenameForOsArch(parsedDatabaseUrl *url.URL) (string, error) {
	result := parsedDatabaseUrl.Path
	if strings.HasPrefix(result, "/") {
		fmt.Printf(">>>>>>>>>>>>>> before: %s", result)
		result = result[2:]
		fmt.Printf(">>>>>>>>>>>>>>  after: %s", result)
	}
	return result, nil
}
