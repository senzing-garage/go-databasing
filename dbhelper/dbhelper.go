package dbhelper

import (
	"fmt"
	"net/url"
	"strings"
)

func ExtractSqliteDatabaseFilename(databaseUrl string) (string, error) {
	var result = ""

	if !strings.HasPrefix(databaseUrl, "sqlite3:") {
		return result, fmt.Errorf("sqlite3 URL schema needed")
	}

	parsedUrl, err := url.Parse(databaseUrl)
	if err != nil {
		return result, err
	}

	if parsedUrl.Scheme != "sqlite3" {
		return result, fmt.Errorf("sqlite3 URL schema needed")
	}

	return extractSqliteDatabaseFilenameForOsArch(parsedUrl)
}
