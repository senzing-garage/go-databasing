package connector

import (
	"context"
	"database/sql/driver"
	"fmt"
	"net"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/senzing-garage/go-databasing/connectormssql"
	"github.com/senzing-garage/go-databasing/connectormysql"
	"github.com/senzing-garage/go-databasing/connectororacle"
	"github.com/senzing-garage/go-databasing/connectorpostgresql"
	"github.com/senzing-garage/go-databasing/connectorsqlite"
	"github.com/senzing-garage/go-helpers/wraperror"
)

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
Function NewConnector is a factory for producing the correct driver.Connector for a given database URL.

Input
  - databaseURL: A properly formed URL containing database connection information
    For acceptable database URLs, see [Database URLs].

[Database URLs]: https://github.com/senzing-garage/go-databasing/blob/main/README.md#database-urls
*/
func NewConnector(ctx context.Context, databaseURL string) (driver.Connector, error) {
	var result driver.Connector

	parsedURL, err := url.Parse(databaseURL)
	if err != nil {
		return result, wraperror.Errorf(err, "connector.NewConnector.urlParse error: %w", err)
	}

	scheme := parsedURL.Scheme
	switch scheme {
	case "sqlite3":
		return createSqlite3Connector(ctx, parsedURL)
	case "postgresql":
		return createPostgresqlConnector(ctx, parsedURL)
	case "mysql":
		return createMysqlConnector(ctx, parsedURL)
	case "mssql":
		return createMssqlConnector(ctx, parsedURL)
	case "oci":
		return createOciConnector(ctx, parsedURL)
	default:
		err = wraperror.Errorf(errPackage, "unknown database scheme: %s error: %w", scheme, errPackage)
	}

	return result, wraperror.Errorf(err, "connector.NewConnector error: %w", err)
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func createSqlite3Connector(ctx context.Context, parsedURL *url.URL) (driver.Connector, error) {
	configuration := parsedURL.Path
	if len(parsedURL.RawQuery) > 0 {
		configuration = fmt.Sprintf("file:%s?%s", configuration[1:], parsedURL.Query().Encode())
	}

	result, err := connectorsqlite.NewConnector(ctx, configuration)

	return result, wraperror.Errorf(err, "connector.createSqlite3Connector error: %w", err)
}

func createPostgresqlConnector(ctx context.Context, parsedURL *url.URL) (driver.Connector, error) {
	// Parse URL: https://pkg.go.dev/net/url#URL
	username := parsedURL.User.Username()
	password, isPasswordSet := parsedURL.User.Password()
	path := parsedURL.Path
	host, port, _ := net.SplitHostPort(parsedURL.Host)

	query, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return nil, wraperror.Errorf(err, "connector.createPostgresqlConnector.url.ParseQuery error: %w", err)
	}

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

	result, err := connectorpostgresql.NewConnector(ctx, configuration)

	return result, wraperror.Errorf(err, "connector.createPostgresqlConnector error: %w", err)
}

func createMysqlConnector(ctx context.Context, parsedURL *url.URL) (driver.Connector, error) {
	// Parse URL: https://pkg.go.dev/net/url#URL
	username := parsedURL.User.Username()
	password, isPasswordSet := parsedURL.User.Password()
	path := parsedURL.Path
	host, port, _ := net.SplitHostPort(parsedURL.Host)

	query, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return nil, wraperror.Errorf(err, "connector.createMysqlConnector.url.ParseQuery error: %w", err)
	}

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
		configuration.Addr += ":" + port
	}

	if len(path) > 1 {
		dbname := strings.ReplaceAll(path, "/", "")
		configuration.DBName = dbname
	} else if schema, ok := query["schema"]; ok {
		configuration.DBName = schema[0]
	}

	result, err := connectormysql.NewConnector(ctx, configuration)

	return result, wraperror.Errorf(err, "connector.createMysqlConnector error: %w", err)
}

func createMssqlConnector(ctx context.Context, parsedURL *url.URL) (driver.Connector, error) {
	// Parse URL: https://pkg.go.dev/net/url#URL
	username := parsedURL.User.Username()
	password, isPasswordSet := parsedURL.User.Password()
	path := parsedURL.Path
	host, port, _ := net.SplitHostPort(parsedURL.Host)

	query, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return nil, wraperror.Errorf(err, "connector.createMssqlConnector.url.ParseQuery error: %w", err)
	}

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

	value, ok := configurationMap["TrustServerCertificate"]
	if ok {
		if value == "yes" {
			configurationMap["TrustServerCertificate"] = "true"
		}
	}

	configuration := ""
	for key, value := range configurationMap {
		configuration += fmt.Sprintf("%s=%s;", key, value)
	}

	result, err := connectormssql.NewConnector(ctx, configuration)

	return result, wraperror.Errorf(err, "connector.createMssqlConnector error: %w", err)
}

func createOciConnector(ctx context.Context, parsedURL *url.URL) (driver.Connector, error) {
	// Parse URL: https://pkg.go.dev/net/url#URL
	username := parsedURL.User.Username()
	password, isPasswordSet := parsedURL.User.Password()
	path := parsedURL.Path
	host, port, _ := net.SplitHostPort(parsedURL.Host)

	query, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return nil, wraperror.Errorf(err, "connector.createOciConnector.url.ParseQuery error: %w", err)
	}

	// See https://pkg.go.dev/github.com/godror/godror

	configurationMap := map[string]string{}
	if len(username) > 0 {
		configurationMap["user"] = username
	}

	if isPasswordSet {
		configurationMap["password"] = password
	}

	configurationMap["connectString"] = fmt.Sprintf("%s:%s%s", host, port, filepath.Clean(path))
	for key, value := range query {
		configurationMap[key] = value[0]
	}

	configuration := ""
	for key, value := range configurationMap {
		configuration += fmt.Sprintf("%s=%s ", key, value)
	}

	result, err := connectororacle.NewConnector(ctx, configuration)

	return result, wraperror.Errorf(err, "connector.createOciConnector error: %w", err)
}
