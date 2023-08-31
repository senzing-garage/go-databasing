package dbhelper

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseMetadata struct {
	databaseFilename string
	databaseUrl      string
	name             string
	succeeds         bool
}

var testCasesForMultiPlatform = []testCaseMetadata{
	{
		name:        "db2-001",
		databaseUrl: "db2://username:password@hostname:50000/G2",
		succeeds:    false,
	},
	{
		name:        "db2-002",
		databaseUrl: "db2://username:password@hostname:50000/G2/?schema=schemaname",
		succeeds:    false,
	},
	{
		name:        "oci-001",
		databaseUrl: "oci://username:password@hostname:1521/G2",
		succeeds:    false,
	},
	{
		name:        "mssql-001",
		databaseUrl: "mssql://username:password@hostname:1433/G2",
		succeeds:    false,
	},
	{
		name:        "mysql-001",
		databaseUrl: "mysql://username:password@hostname:3306/G2",
		succeeds:    false,
	},
	{
		name:        "oci-001",
		databaseUrl: "oci://username:password@hostname:1521/G2",
		succeeds:    false,
	},
	{
		name:        "postgresql-001",
		databaseUrl: "postgresql://username:password@hostname:5432/G2",
		succeeds:    false,
	},
	{
		name:        "postgresql-002",
		databaseUrl: "postgresql://username:password@hostname:5432/G2/?schema=schemaname",
		succeeds:    false,
	},
}

var testCases = append(testCasesForMultiPlatform, testCasesForOsArch...)

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error = nil
	return err
}

func teardown() error {
	var err error = nil
	return err
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func testError(test *testing.T, err error, succeeds bool) {
	if succeeds {
		if err != nil {
			assert.FailNow(test, err.Error())
		}
	} else {
		if err == nil {
			assert.FailNow(test, "failure was expected")
		}
	}

}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestExtractSqliteDatabaseFilename(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			result, err := ExtractSqliteDatabaseFilename(testCase.databaseUrl)
			testError(test, err, testCase.succeeds)
			if len(testCase.databaseFilename) > 0 {
				assert.Equal(test, testCase.databaseFilename, result)
			}
		})
	}
}
