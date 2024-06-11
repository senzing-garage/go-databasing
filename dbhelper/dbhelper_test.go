package dbhelper

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseMetadata struct {
	databaseFilename string
	databaseURL      string
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
}

var testCases = append(testCasesForMultiPlatform, testCasesForOsArch...)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestExtractSqliteDatabaseFilename(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			result, err := ExtractSqliteDatabaseFilename(testCase.databaseURL)
			testError(test, err, testCase.succeeds)
			if len(testCase.databaseFilename) > 0 {
				assert.Equal(test, testCase.databaseFilename, result)
			}
		})
	}
}

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
	var err error
	return err
}

func teardown() error {
	var err error
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
