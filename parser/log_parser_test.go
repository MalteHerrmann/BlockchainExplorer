package logparser

import (
	"testing"

	"github.com/MalteHerrmann/BlockchainExplorer/testutil"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
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
			require.Equal(t, lp.sub, nil, "sub should be initialized")
			require.Equal(t, lp.lastBlockNumber, uint64(0), "lastBlockNumber should be 0")
			require.Equal(t, lp.lastTxIndex, uint(0), "lastTxIndex should be 0")
			require.Equal(t, lp.lastTxHash, common.Hash{}, "lastTxHash should be zero-value for hash type")
			require.Equal(t, lp.nTxInLastBlock, 0, "nTxInLastBlock should be 0")
			require.Equal(t, lp.GetLastBlockNumber(), uint64(0), "lastBlockNumber should be 0")
			require.Equal(t, lp.GetLastTxIndex(), uint(0), "lastTxIndex should be 0")
			require.Equal(t, lp.GetLastTxHash(), common.Hash{}, "lastTxHash should be zero-value for hash type")
			require.Equal(t, lp.GetNumberOfTransactionsInLastBlock(), 0, "nTxInLastBlock should be 0")
		})
	}
}

func TestNewLogParserWithURL(t *testing.T) {
	// Define testcases
	testcases := []struct {
		name    string
		url     string
		expPass bool
	}{
		{
			"url is empty",
			"",
			false,
		},
		{
			"url is invalid",
			"invalid url",
			false,
		},
		{
			"url is valid",
			"http://localhost:8545",
			true,
		},
	}

	// Run tests
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			lp, err := NewLogParserWithURL(tc.url)
			if tc.expPass {
				require.NoError(t, err, "NewLogParserWithURL should not return an error")
				require.NotNil(t, lp, "NewLogParserWithURL should return a LogParser")
			} else {
				require.Error(t, err, "NewLogParserWithURL should return an error")
				require.Nil(t, lp, "NewLogParserWithURL should return a LogParser")
			}
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

	// Process a log with a new empty transaction
	lp.processLog(ethtypes.Log{
		BlockNumber: 1,
		TxIndex:     2,
		TxHash:      common.Hash{},
	})
	require.Equal(t, lp.lastBlockNumber, uint64(1), "lastBlockNumber should be 1")
	require.Equal(t, lp.lastTxIndex, uint(2), "lastTxIndex should be 2")
	require.Equal(t, lp.lastTxHash, common.Hash{}, "lastTxHash should be zero-value for hash type")
}
