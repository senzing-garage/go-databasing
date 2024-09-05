package connector

import (
	"context"
	"database/sql/driver"
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/senzing-garage/go-databasing/connectormssql"
	"github.com/senzing-garage/go-databasing/connectormysql"
	"github.com/senzing-garage/go-databasing/connectorpostgresql"
	"github.com/senzing-garage/go-databasing/connectorsqlite"
)

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
Function NewConnector is a factory for producing the correct driver.Connector for a given database URL.

Input
  - databaseUrl: A properly formed URL containing database connection information
    For acceptable database URLs, see [Database URLs].

[Database URLs]: https://github.com/senzing-garage/go-databasing/blob/main/README.md#use
*/
func NewConnector(ctx context.Context, databaseURL string) (driver.Connector, error) {
	var result driver.Connector

	parsedURL, err := url.Parse(databaseURL)
	if err != nil {

		if strings.HasPrefix(databaseURL, "postgresql") {
			index := strings.LastIndex(databaseURL, ":")
			newDatabaseURL := databaseURL[:index] + "/" + databaseURL[index+1:]
			parsedURL, err = url.Parse(newDatabaseURL)
		}

		if err != nil {
			return result, err
		}
	}

	// Parse URL: https://pkg.go.dev/net/url#URL

	scheme := parsedURL.Scheme
	username := parsedURL.User.Username()
	password, isPasswordSet := parsedURL.User.Password()
	path := parsedURL.Path
	// fragment := parsedUrl.Fragment
	host, port, _ := net.SplitHostPort(parsedURL.Host)
	query, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return result, err
	}

	switch scheme {
	case "sqlite3":
		configuration := path
		result, err = connectorsqlite.NewConnector(ctx, configuration)

	case "postgresql":
		// See https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
		configurationMap := map[string]string{}
		if len(username) > 0 {
			configurationMap["user"] = username
		}
		if isPasswordSet {
			configurationMap["password"] = password
		}
		if len(host) > 0 {
			configurationMap["host"] = host
		}
		if len(port) > 0 {
			configurationMap["port"] = port
		}
		if len(path) > 1 {
			dbname := strings.ReplaceAll(path, "/", "")
			configurationMap["dbname"] = dbname
		}
		for key, value := range query {
			configurationMap[key] = value[0]
		}
		if searchPath, ok := query["schema"]; ok {
			configurationMap["search_path"] = searchPath[0]
		}
		configuration := ""
		for key, value := range configurationMap {
			configuration += fmt.Sprintf("%s=%s ", key, value)
		}
		result, err = connectorpostgresql.NewConnector(ctx, configuration)

	case "mysql":
		// See https://pkg.go.dev/github.com/go-sql-driver/mysql#Confi
		configuration := &mysql.Config{
			Net:       "tcp",
			Collation: "utf8mb4_general_ci",
		}
		if len(username) > 0 {
			configuration.User = username
		}
		if isPasswordSet {
			configuration.Passwd = password
		}
		if len(host) > 0 {
			configuration.Addr = host
		}
		if len(port) > 0 {
			configuration.Addr += fmt.Sprintf(":%s", port)
		}
		if len(path) > 1 {
			dbname := strings.ReplaceAll(path, "/", "")
			configuration.DBName = dbname
		} else if schema, ok := query["schema"]; ok {
			configuration.DBName = schema[0]
		}

		result, err = connectormysql.NewConnector(ctx, configuration)

	case "mssql":
		// See https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn
		// databaseConnector, err = connectormssql.NewConnector(ctx, "user id=sa;password=Passw0rd;database=master;server=localhost")
		configurationMap := map[string]string{}
		if len(username) > 0 {
			configurationMap["user id"] = username
		}
		if isPasswordSet {
			configurationMap["password"] = password
		}
		if len(host) > 0 {
			configurationMap["server"] = host
		}
		if len(port) > 0 {
			configurationMap["port"] = port
		}
		if len(path) > 1 {
			dbname := strings.ReplaceAll(path, "/", "")
			configurationMap["database"] = dbname
		}
		for key, value := range query {
			configurationMap[key] = value[0]
		}
		configuration := ""
		for key, value := range configurationMap {
			configuration += fmt.Sprintf("%s=%s;", key, value)
		}
		result, err = connectormssql.NewConnector(ctx, configuration)

	default:
		err = fmt.Errorf("unknown database scheme: %s", scheme)
	}

	return result, err
}
