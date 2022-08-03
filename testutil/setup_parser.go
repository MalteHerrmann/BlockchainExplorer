package testutil

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// SetupClientForTesting creates a new ethereum client, that
// is preconfigured for testing purposes.
func SetupClientForTesting() (*ethclient.Client, error) {
	// Specify the URL to connect to
	url := "wss://mainnet.infura.io/ws/v3/266df44dd6b24b39ba5bb703049aa3c8"
	//url := "wss://eth.bd.evmos.org:8546"

	// Create a new instance of the API client from the client package.
	// The client is used to communicate with the node.
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return client, nil
}
