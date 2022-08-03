package logparser

import (
	"github.com/MalteHerrmann/TrackGovernance/testutil"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewLogParser(t *testing.T) {
	// Prepare tests
	connectedClient, err := testutil.SetupClientForTesting()
	require.NoError(t, err, "Client for testing should connect without error.")

	// Define testcases
	testcases := []struct {
		name    string
		client  *ethclient.Client
		expPass bool
	}{
		{
			"client is nil",
			nil,
			false,
		},
		{
			"client is connected",
			connectedClient,
			false,
		},
	}

	// Run tests
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			lp := NewLogParser(tc.client)
			require.Equal(t, lp.client, tc.client, "client should be set")
			require.Equal(t, lp.logs, make(chan ethtypes.Log), "logs should be initialized")
			require.Equal(t, lp.sub, nil, "sub should be initialized")
			require.Equal(t, lp.lastBlockNumber, uint64(0), "lastBlockNumber should be 0")
			require.Equal(t, lp.lastTxIndex, uint(0), "lastTxIndex should be 0")
			require.Equal(t, lp.lastTxHash, common.Hash{}, "lastTxHash should be zero-value for hash type")
			require.Equal(t, lp.FailedTxs, []common.Hash{}, "FailedTxs should be initialized")
		})
	}
}

func TestProcessLog(t *testing.T) {
	// Prepare tests
	client, err := testutil.SetupClientForTesting()
	require.NoError(t, err, "Client should be created without errors")

	lp := NewLogParser(client)

	// Process log
	lp.processLog(ethtypes.Log{})
	require.Equal(t, lp.lastBlockNumber, uint64(0), "lastBlockNumber should be 0")
	require.Equal(t, lp.lastTxIndex, uint(0), "lastTxIndex should be 0")
	require.Equal(t, lp.lastTxHash, common.Hash{}, "lastTxHash should be zero-value for hash type")
}
