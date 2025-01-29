package main

import (
	"lichess-tui/internal/config"
	"lichess-tui/internal/requests"
	"lichess-tui/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Load the config, if it doesn't exist, create it
	config.LoadConfig(config.GetConfigPath())

	p := tea.NewProgram(tui.NewModel())

	cfg := config.GetConfig()
	go requests.StreamIncomingEvents(cfg.Token, p)

	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
