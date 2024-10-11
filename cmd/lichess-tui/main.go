package main

import (
	"github.com/rivo/tview"
	// "fmt"
)

type Piece uint8

const (
	NAME      string = "Lichess TUI"
	CELL_SIZE int    = 1
)

const (
	Pawn   Piece = 0
	Rook   Piece = 1
	Knight Piece = 2
	Bishop Piece = 3
	Queen  Piece = 4
	King   Piece = 5
)

func newPrimitive(text string) tview.Primitive {
	return tview.NewBox().SetBorder(true)
}
func main() {
	// menu := newPrimitive("Menu")
	// main := newPrimitive("Main content")
	// sideBar := newPrimitive("Side Bar")
	//
	grid := tview.NewGrid().
		SetColumns(8).SetRows(8).
		SetBorders(true)
		// AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
		// AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	for i := range 8 {
		for j := range 8 {
			grid.AddItem(newPrimitive("0"), i, j, 1, 1, 1, 1, false)
		}
	}

	if err := tview.NewApplication().SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
