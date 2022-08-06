package explorergui

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestNewExplorerGui tests if an ExplorerGUI struct can
// be created.
func TestNewExplorerGUI(t *testing.T) {
	pg := NewExplorerGUI()
	require.NotNil(t, pg, "ExplorerGui struct could not be created.")
}
