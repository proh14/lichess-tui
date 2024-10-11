package ui

import (
	"github.com/proh14/lichess-tui/internal/ui/components"
	"github.com/rivo/tview"
)

type Ui struct {
	Board *tview.Grid
	App   *tview.Application
}

var UI = Ui{Board: nil, App: nil}

func InitUI() {
	setStyle()
	UI.Board = components.InitBoard()
	UI.App = tview.NewApplication().SetRoot(UI.Board, true).EnableMouse(true)
}
