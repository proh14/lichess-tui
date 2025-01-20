package main

import (
	"lichess-tui/internal/config"
	"lichess-tui/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Load the config, if it doesn't exist, create it
	config.LoadConfig(config.GetConfigPath())

	p := tea.NewProgram(tui.NewModel(), tea.WithAltScreen())

	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
