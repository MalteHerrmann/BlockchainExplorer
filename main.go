// TrackGovernance can be used to inspect governance information
// on Evmos.
package main

import (
	"fmt"
	logparser "github.com/MalteHerrmann/TrackGovernance/parser"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	// Print info
	fmt.Println("")
	fmt.Println("----------------------------------------------------")
	fmt.Println("                 TrackGovernance")
	fmt.Println("----------------------------------------------------")
	// Create a new instance of the API client from the client package.
	// The client is used to communicate with the node.
	//
	url := "wss://mainnet.infura.io/ws/v3/266df44dd6b24b39ba5bb703049aa3c8"
	//url := "wss://eth.bd.evmos.org:8546"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	fmt.Println("Connected to Ethereum client at " + url)
	fmt.Println("")

	// Create log parser
	lp := logparser.NewLogParser(client)

	// Subscribe to new blocks
	lp.SubscribeToBlocks()
}
