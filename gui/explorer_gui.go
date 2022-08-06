// Package explorergui contains the user interface for the blockchain
// explorer.
package explorergui

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	logparser "github.com/MalteHerrmann/TrackGovernance/parser"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ExplorerGUI is the type for the user interface to
// display information about the connected blockchain.
type ExplorerGUI struct {
	// The app which is used to create the Window.
	app fyne.App

	// The Window which is used to display the content.
	Window fyne.Window

	// The client used to communicate with the node.
	client *ethclient.Client

	// The logparser used to track information
	lp *logparser.LogParser
}

// NewExplorerGUI creates a new ExplorerGUI.
func NewExplorerGUI() *ExplorerGUI {
	eg := new(ExplorerGUI)

	eg.app = app.New()
	eg.Window = eg.app.NewWindow("Blockchain Parser")

	// Set Window size
	eg.Window.Resize(fyne.NewSize(800, 600))

	// Add label with current block height
	label := widget.NewLabel("Waiting for info...")

	// Assign content to window
	eg.Window.SetContent(label)

	return eg
}

// Connect creates a client, which connects to the defined
// url and starts the parser.
func (eg *ExplorerGUI) Connect(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	eg.client = client

	return eg.client, nil
}

// GetParser returns the parser used to track information.
func (eg *ExplorerGUI) GetParser() *logparser.LogParser {
	eg.lp = logparser.NewLogParser(eg.client)

	return eg.lp
}

// GetBlockNumber returns the current height of the blockchain.
func (eg *ExplorerGUI) GetBlockNumber() uint64 {
	blocknumber, err := eg.client.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Blocknumber could not be retrieved.")
	}

	return blocknumber
}

// UpdateGUI updates the label with the current block height.
func (eg *ExplorerGUI) UpdateGUI() {
	heightLabel := widget.NewLabel("Current blocknumber: " + fmt.Sprintf("%d", eg.lp.LastBlockNumber))
	txIndexLabel := widget.NewLabel("Current tx index: " + fmt.Sprintf("%v", eg.lp.LastTxIndex))
	txHashLabel := widget.NewLabel("Current tx index: " + fmt.Sprintf("%v", eg.lp.LastTxHash))

	// Create new v box to display the content
	vbox := container.New(layout.NewVBoxLayout(), heightLabel, txIndexLabel, txHashLabel)

	eg.Window.SetContent(vbox)
}
