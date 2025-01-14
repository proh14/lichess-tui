package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/proh14/lichess-tui/internal/config"
	"github.com/proh14/lichess-tui/internal/tui"
)

func main() {
	// Load the config, if it doesn't exist, create it
	config.LoadConfig(config.GetConfigPath())

	p := tea.NewProgram(tui.NewModel())

	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
