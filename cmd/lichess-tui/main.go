package main

import (
	// "lichess-tui/internal/config"
	// "lichess-tui/internal/tui"
	"lichess-tui/internal/requests"

	// tea "github.com/charmbracelet/bubbletea"
)

func main() {
	respVar := requests.SeekGameResponse{}
	// cfg := config.GetConfig()
	requests.SeekGame(requests.SeekGameConfig{Time: 10, Increment: 0}, "lip_NUvDrFiwfjsW0tneprnA", &respVar)

	// // Load the config, if it doesn't exist, create it
	// config.LoadConfig(config.GetConfigPath())
	//
	// p := tea.NewProgram(tui.NewModel(), tea.WithAltScreen())
	//
	// _, err := p.Run()
	// if err != nil {
	// 	panic(err)
	// }
}
