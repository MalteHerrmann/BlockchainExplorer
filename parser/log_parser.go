package logparser

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// LogParser is a type that can be used to parse logs.
type LogParser struct {
	// The client used to communicate with the node.
	client *ethclient.Client

	// The channel used to receive new blocks.
	logs chan ethtypes.Log

	// The subscription used to receive new blocks.
	sub ethereum.Subscription

	// Last seen block number
	lastBlockNumber uint64

	// Last seen transaction info
	lastTxIndex uint
	lastTxHash  common.Hash

	// Failed transactions
	FailedTxs []common.Hash
}

// NewLogParser creates a new LogParser.
func NewLogParser(client *ethclient.Client) *LogParser {
	return &LogParser{
		client: client,
		logs:   make(chan ethtypes.Log),
	}
}

// SubscribeToBlocks subscribes to new blocks and processes the
// received information.
func (lp *LogParser) SubscribeToBlocks() {
	// Subscribe with empty filter settings to receive all logs
	sub, err := lp.client.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{}, lp.logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}
	lp.sub = sub

	for {
		select {
		case err := <-lp.sub.Err():
			log.Fatalf("Failed to subscribe to logs: %v", err)
		case block := <-lp.logs:
			lp.processLog(block)
		}
	}
}

// ProcessLog is responsible for the further processing of a received
// log.
func (lp *LogParser) processLog(block ethtypes.Log) {
	currentBlocknumber := block.BlockNumber
	currentTxIndex := block.TxIndex
	fmt.Printf("\rBlock %d - tx %4d", currentBlocknumber, currentTxIndex)

	if currentBlocknumber > lp.lastBlockNumber {
		lp.lastBlockNumber = currentBlocknumber
	}
	if currentTxIndex > lp.lastTxIndex {
		// always processes the last tx when a new tx comes in
		ok := lp.ProcessTx(lp.lastTxHash)
		if !ok {
			lp.FailedTxs = append(lp.FailedTxs, lp.lastTxHash)
			fmt.Printf(" Failed transactions: %v", len(lp.FailedTxs))
		}

		// Assign new tx info
		lp.lastTxIndex = block.TxIndex
		lp.lastTxHash = block.TxHash
	}

	fmt.Printf("") // TODO: remove this line - somehow tests are terminated without it, why?
}

// ProcessTx queries the transaction with the given
// transaction hash and processes the received information.
func (lp *LogParser) ProcessTx(txHash common.Hash) bool {
	_, isPending, err := lp.client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		fmt.Printf(" Failed to get transaction: %v", txHash.Hex())
		return false
	}

	if isPending {
		fmt.Printf(" Transaction %v is pending", txHash.Hex())
	} else {
		fmt.Printf(" Transaction %v is not pending", txHash.Hex())
	}

	return true
}
