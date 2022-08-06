package explorergui

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	URL = "wss://mainnet.infura.io/ws/v3/266df44dd6b24b39ba5bb703049aa3c8"
)

func TestNewExplorerGUI(t *testing.T) {
	pg := NewExplorerGUI()
	require.NotNil(t, pg, "ExplorerGui struct could not be created.")
}

func TestConnect(t *testing.T) {
	pg := NewExplorerGUI()
	_, err := pg.Connect(URL)
	require.NoError(t, err, "Failed to connect to the Ethereum client.")
}

func TestUpdateBlockNumber(t *testing.T) {
	pg := NewExplorerGUI()
	_, err := pg.Connect(URL)
	require.NoError(t, err, "Failed to connect to the Ethereum client.")
	pg.UpdateBlockNumber()
}

func TestGetBlockNumber(t *testing.T) {
	pg := NewExplorerGUI()
	_, err := pg.Connect(URL)
	require.NoError(t, err, "Failed to connect to the Ethereum client.")
	blocknumber := pg.GetBlockNumber()
	require.NotNil(t, blocknumber, "Blocknumber could not be retrieved.")
}
