// Package explorergui contains the user interface for the blockchain
// explorer.
package explorergui

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
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
}

// NewExplorerGUI creates a new ExplorerGUI.
func NewExplorerGUI() *ExplorerGUI {
	eg := new(ExplorerGUI)

	eg.app = app.New()
	eg.Window = eg.app.NewWindow("Blockchain Parser")

	// Set Window size
	eg.Window.Resize(fyne.NewSize(800, 600))

	// Add label with current block height
	label := widget.NewLabel("Current blocknumber: ")

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

// GetBlockNumber returns the current height of the blockchain.
func (eg *ExplorerGUI) GetBlockNumber() uint64 {
	blocknumber, err := eg.client.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Blocknumber could not be retrieved.")
	}

	return blocknumber
}

// UpdateBlockNumber updates the label with the current block height.
func (eg *ExplorerGUI) UpdateBlockNumber() {
	label := widget.NewLabel("Current blocknumber: " + fmt.Sprintf("%d", eg.GetBlockNumber()))
	eg.Window.SetContent(label)
}
