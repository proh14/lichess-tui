package components

import (
	"github.com/proh14/lichess-tui/internal/config"
	"github.com/rivo/tview"
)

func newPrimitive(text string) tview.Primitive {
	return tview.NewButton(text)
}

func InitBoard() *tview.Grid {
	board := make([][]*tview.Grid, 8)

	for i := range 8 {
		board[i] = make([]*tview.Grid, 8)
	}

	grid := tview.NewGrid().
		SetSize(8, 8, 2*config.CELL_SIZE, 8*config.CELL_SIZE).
		SetBorders(true)
		// AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
		// AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	for i := range 8 {
		for j := range 8 {
			board[i][j] = grid.AddItem(newPrimitive(config.PawnStr), i, j, 1, 1, 0, 0, false)
		}
	}

	grid.AddItem(newPrimitive("Surrender"), 8, 0, 1, 8, 0, 0, false)
	grid.AddItem(newPrimitive("Offer draw"), 9, 0, 1, 8, 0, 0, false)
	grid.AddItem(newPrimitive("Propose takeback"), 10, 0, 1, 8, 0, 0, false)

	return grid
}
