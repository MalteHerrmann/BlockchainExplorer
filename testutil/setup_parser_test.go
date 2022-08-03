package testutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetupClientForTesting(t *testing.T) {
	_, err := SetupClientForTesting()
	require.NoError(t, err, "Client should be created without errors")
}
