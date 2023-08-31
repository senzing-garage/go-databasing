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
		fmt.Printf(">>>>>>>>>>>>>> before: %s\n", result)
		result = result[1:]
		fmt.Printf(">>>>>>>>>>>>>>  after: %s\n", result)
	}
	return result, nil
}
