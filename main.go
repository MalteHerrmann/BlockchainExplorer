// TrackGovernance can be used to inspect governance information
// on Evmos.
package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// SubscribeToBlocks subscribes to new blocks and prints them to the console.
func SubscribeToBlocks(client *ethclient.Client) {
	// Create a subscription that fires on new blocks.
	ctx := context.Background()

	logs := make(chan ethtypes.Log)
	sub, err := client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{}, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}
	fmt.Println("Subscribed to logs")

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Failed to subscribe to logs: %v", err)
		case block := <-logs:
			fmt.Println("Block: ", block.BlockHash.Hex())
		}
	}
}

func main() {
	// Create a new instance of the API client from the client package.
	// The client is used to communicate with the node.
	//
	//client, err := ethclient.Dial("wss://eth.bd.evmos.org:8546")
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/266df44dd6b24b39ba5bb703049aa3c8")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Subscribe to new blocks
	SubscribeToBlocks(client)
}
