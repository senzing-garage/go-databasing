package connector

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCaseMetadata struct {
	databaseURL      string
	isBadDatabaseURL bool
	name             string
}

var testCases = []testCaseMetadata{
	{
		name:        "oci-001",
		databaseURL: "oci://username:password@hostname:1521/G2",
	},
	{
		name:        "oci-002",
		databaseURL: "oci://sys:Passw0rd@localhost:1521/FREE/?sysdba=true&noTimezoneCheck=true",
	},
	{
		name:             "oci-003",
		databaseURL:      "oci://username:password@G2",
		isBadDatabaseURL: true,
	},
	{
		name:        "mssql-001",
		databaseURL: "mssql://username:password@hostname:1433/G2",
	},
	{
		name:        "mssql-002",
		databaseURL: "mssql://username:password@hostname:1433/G2/?driver=mssqldriver",
	},
	{
		name:        "mssql-003",
		databaseURL: "mysql://username:password@hostname:3306/?schema=schemaname",
	},
	{
		name:        "mssql-004",
		databaseURL: "mssql://username:password@hostname:3306/database?schema=schemaname",
	},
	{
		name:             "mssql-005",
		databaseURL:      "mssql://username:password@hostname:1433:G2/?driver=mssqldriver",
		isBadDatabaseURL: true,
	},
	{
		name:             "mssql-006",
		databaseURL:      "mssql://username:password@G2",
		isBadDatabaseURL: true,
	},
	{
		name:        "mysql-001",
		databaseURL: "mysql://username:password@hostname:3306/G2",
	},
	{
		name:        "mysql-002",
		databaseURL: "mysql://username:password@hostname:3306/database?schema=schemaname",
	},
	{
		name:             "mysql-003",
		databaseURL:      "mysql://username:password@hostname:3306/?schema=G2",
		isBadDatabaseURL: true,
	},
	{
		name:        "postgresql-001",
		databaseURL: "postgresql://username:password@hostname:5432/G2",
	},
	{
		name:        "postgresql-002",
		databaseURL: "postgresql://username:password@hostname:5432/G2/?schema=schemaname",
	},
	{
		name:        "postgresql-003",
		databaseURL: "postgresql://username:password@hostname:5432:G2/?schema=schemaname",
	},
	{
		name:        "postgresql-004",
		databaseURL: "postgresql://username:password@hostname:5432/database/?schema=schemaname",
	},
	{
		name:             "postgresql-005",
		databaseURL:      "postgresql://username:password@hostname:5432:database/?schema=schemaname",
		isBadDatabaseURL: true,
	},
	{
		name:             "postgresql-006",
		databaseURL:      "postgresql://username:password@hostname:5432:G2/",
		isBadDatabaseURL: true,
	},
	{
		name:        "sqlite-001",
		databaseURL: "sqlite3://na:na@/tmp/sqlite/G2C.db",
	},
	{
		name:             "bad-protocol-001",
		databaseURL:      "badProtocol://username:password@hostname:3306/database?schema=schemaname",
		isBadDatabaseURL: true,
	},
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestNewConnector(test *testing.T) {
	ctx := context.TODO()
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			_, err := NewConnector(ctx, testCase.databaseURL)
			if testCase.isBadDatabaseURL {
				require.Error(test, err)
			} else {
				require.NoError(test, err)
			}
		})
	}
}
