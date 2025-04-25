package dbhelper_test

import (
	"testing"

	"github.com/senzing-garage/go-databasing/dbhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCaseMetadata struct {
	databaseFilename string
	databaseURL      string
	fixedDatabaseURL string
	name             string
	succeeds         bool
}

var testCasesForMultiPlatform = []testCaseMetadata{
	{
		name:        "db2-001",
		databaseURL: "db2://username:password@hostname:50000/G2",
		succeeds:    false,
	},
	{
		name:        "db2-002",
		databaseURL: "db2://username:password@hostname:50000/G2/?schema=schemaname",
		succeeds:    false,
	},
	{
		name:        "oci-001",
		databaseURL: "oci://username:password@hostname:1521/G2",
		succeeds:    false,
	},
	{
		name:        "mssql-001",
		databaseURL: "mssql://username:password@hostname:1433/G2",
		succeeds:    false,
	},
	{
		name:        "mysql-001",
		databaseURL: "mysql://username:password@hostname:3306/G2",
		succeeds:    false,
	},
	{
		name:        "oci-001",
		databaseURL: "oci://username:password@hostname:1521/G2",
		succeeds:    false,
	},
	{
		name:        "postgresql-001",
		databaseURL: "postgresql://username:password@hostname:5432/G2",
		succeeds:    false,
	},
	{
		name:        "postgresql-002",
		databaseURL: "postgresql://username:password@hostname:5432/G2/?schema=schemaname",
		succeeds:    false,
	},
	{
		name:             "postgresql-003",
		databaseURL:      "postgresql://username:password@hostname:5432:G2/?schema=schemaname",
		fixedDatabaseURL: "postgresql://username:password@hostname:5432/G2/?schema=schemaname",
		succeeds:         false,
	},
	{
		name:        "sqlite-001",
		databaseURL: "sqlite3://na:na@/tmp/sqlite/G2C.db",
		succeeds:    true,
	},
}

var testCases = append(testCasesForMultiPlatform, testCasesForOsArch...)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestExtractSqliteDatabaseFilename(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			result, err := dbhelper.ExtractSqliteDatabaseFilename(testCase.databaseURL)
			if testCase.succeeds {
				require.NoError(test, err)
			} else {
				require.Error(test, err)
			}

			if len(testCase.databaseFilename) > 0 {
				assert.Equal(test, testCase.databaseFilename, result)
			}
		})
	}
}

func TestGetMessenger(test *testing.T) {
	test.Parallel()

	options := []interface{}{}
	_ = dbhelper.GetMessenger(1, map[int]string{}, 0, options...)
}
