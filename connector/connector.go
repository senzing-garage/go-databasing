package connector

import (
	"context"
	"database/sql/driver"
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/senzing/go-databasing/connectormssql"
	"github.com/senzing/go-databasing/connectormysql"
	"github.com/senzing/go-databasing/connectorpostgresql"
	"github.com/senzing/go-databasing/connectorsqlite"
)

// ----------------------------------------------------------------------------
// Constructor methods
// ----------------------------------------------------------------------------

/*
A factory for producing the correct driver.Connector for a given database URL.

Input
  - databaseUrl: A properly formed URL containing database connection information
    For acceptable database URLs, see https://....
*/
func NewConnector(ctx context.Context, databaseUrl string) (driver.Connector, error) {
	var result driver.Connector = nil

	parsedUrl, err := url.Parse(databaseUrl)
	if err != nil {

		if strings.HasPrefix(databaseUrl, "postgresql") {
			index := strings.LastIndex(databaseUrl, ":")
			newDatabaseUrl := databaseUrl[:index] + "/" + databaseUrl[index+1:]
			parsedUrl, err = url.Parse(newDatabaseUrl)
		}

		if err != nil {
			return result, err
		}
	}

	// Parse URL: https://pkg.go.dev/net/url#URL

	scheme := parsedUrl.Scheme
	username := parsedUrl.User.Username()
	password, isPasswordSet := parsedUrl.User.Password()
	path := parsedUrl.Path
	// fragment := parsedUrl.Fragment
	host, port, _ := net.SplitHostPort(parsedUrl.Host)
	query, err := url.ParseQuery(parsedUrl.RawQuery)
	if err != nil {
		return result, err
	}

	switch scheme {
	case "sqlite3":
		configuration := path
		fmt.Printf("sqlite3: %s\n", configuration)
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
			dbname := strings.Replace(path, "/", "", -1)
			configurationMap["dbname"] = dbname
		}
		for key, value := range query {
			configurationMap[key] = value[0]
		}
		if search_path, ok := query["schema"]; ok {
			configurationMap["search_path"] = search_path[0]
		}
		configuration := ""
		for key, value := range configurationMap {
			configuration = configuration + fmt.Sprintf("%s=%s ", key, value)
		}
		fmt.Printf("postgresql: %s\n", configuration)
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
			configuration.Addr = configuration.Addr + fmt.Sprintf(":%s", port)
		}
		if len(path) > 1 {
			dbname := strings.Replace(path, "/", "", -1)
			configuration.DBName = dbname
		} else if schema, ok := query["schema"]; ok {
			configuration.DBName = schema[0]
		}

		fmt.Printf("mysql: %v\n", configuration)
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
			dbname := strings.Replace(path, "/", "", -1)
			configurationMap["database"] = dbname
		}
		for key, value := range query {
			configurationMap[key] = value[0]
		}
		configuration := ""
		for key, value := range configurationMap {
			configuration = configuration + fmt.Sprintf("%s=%s;", key, value)
		}
		fmt.Printf("mssql: %v\n", configuration)
		result, err = connectormssql.NewConnector(ctx, configuration)

	default:
		err = fmt.Errorf("unknown database scheme: %s", scheme)
	}

	return result, err
}
