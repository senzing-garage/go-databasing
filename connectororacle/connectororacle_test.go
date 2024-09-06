package connectororacle

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestNewConnector(test *testing.T) {
	ctx := context.TODO()
	// configuration := `user="sysdba" password="Passw0rd" connectString="localhost:1521/FREEPDB1"`
	configuration := `user="sysdba" password="Passw0rd" connectString="localhost:1521/FREE"`
	databaseConnector, err := NewConnector(ctx, configuration)

	for _, e := range os.Environ() {
		fmt.Println(e)
	}

	require.NoError(test, err)
	_, err = databaseConnector.Connect(ctx)
	require.NoError(test, err)
}
