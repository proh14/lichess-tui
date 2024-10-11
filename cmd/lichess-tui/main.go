package main

import (
	"github.com/proh14/lichess-tui/internal/ui"
)

func main() {

	ui.InitUI()

	if err := ui.UI.App.Run(); err != nil {
		panic(err)
	}
}
