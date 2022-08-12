package explorergui

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	URL = "wss://mainnet.infura.io/ws/v3/266df44dd6b24b39ba5bb703049aa3c8"
)

func TestNewExplorerGUI(t *testing.T) {
	eg := NewExplorerGUI()
	require.NotNil(t, eg, "ExplorerGui struct could not be created.")
}

func TestConnect(t *testing.T) {
	eg := NewExplorerGUI()
	_, err := eg.Connect(URL)
	require.NoError(t, err, "Failed to connect to the Ethereum client.")
}

func TestUpdateGUI(t *testing.T) {
	eg := NewExplorerGUI()
	_, err := eg.Connect(URL)
	require.NoError(t, err, "Failed to connect to the Ethereum client.")
	// TODO: implement logparser for this test
	eg.UpdateGUI()
}

func TestGetBlockNumber(t *testing.T) {
	eg := NewExplorerGUI()
	_, err := eg.Connect(URL)
	require.NoError(t, err, "Failed to connect to the Ethereum client.")
	blocknumber := eg.GetBlockNumber()
	require.NotNil(t, blocknumber, "Blocknumber could not be retrieved.")
}
