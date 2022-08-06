// Package parsergui contains the user interface for the blockchain
// parser.
package parsergui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	logparser "github.com/MalteHerrmann/TrackGovernance/parser"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ParserGUI is the type for the user interface to
// display information about the connected blockchain.
type ParserGUI struct {
	// The app which is used to create the window.
	app fyne.App

	// The window which is used to display the content.
	window fyne.Window

	// The client used to communicate with the node.
	client *ethclient.Client

	// The parser used to parse the logs.
	parser *logparser.LogParser
}

// NewParserGUI creates a new ParserGUI.
func NewParserGUI() *ParserGUI {
	pg := new(ParserGUI)

	pg.app = app.New()
	pg.window = pg.app.NewWindow("Blockchain Parser")

	// Set window size
	pg.window.Resize(fyne.NewSize(800, 600))

	return pg
}
