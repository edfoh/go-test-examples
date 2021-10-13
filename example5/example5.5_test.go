// +build all

package example5_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSomething(t *testing.T) {
	envVarCleanup := testSetEnvVar(t)
	defer envVarCleanup()

	// run some tests
}

func testSetEnvVar(t *testing.T) func() {
	t.Helper()
	err := os.Setenv("my_setting", "true")
	require.NoError(t, err)

	return func() { os.Setenv("my_setting", "") }
}
