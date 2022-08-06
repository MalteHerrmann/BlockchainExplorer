// TrackGovernance can be used to inspect governance information
// on Evmos.
package main

import (
	"fmt"
	explorergui "github.com/MalteHerrmann/TrackGovernance/gui"
	"log"
	"time"
)

func main() {
	url := "wss://mainnet.infura.io/ws/v3/266df44dd6b24b39ba5bb703049aa3c8"
	//url := "wss://eth.bd.evmos.org:8546"

	gui := explorergui.NewExplorerGUI()
	_, err := gui.Connect(url)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	fmt.Println("Connected to Ethereum client at " + url)
	fmt.Println("")

	lp := gui.GetParser()
	go func() {
		lp.SubscribeToBlocks()
	}()

	go func() {
		for range time.Tick(time.Second) {
			gui.UpdateBlockNumber()
		}
	}()
	gui.Window.ShowAndRun()
}
