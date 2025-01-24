package quickgame

import (
	"fmt"

	"lichess-tui/internal/config"
	"lichess-tui/internal/requests"
	"lichess-tui/internal/tui/grid"
	"lichess-tui/internal/tui/message"

	tea "github.com/charmbracelet/bubbletea"
)

var timeFormats = [...]message.StartGame{
	{1, 0},
	{2, 1},
	{3, 0},
	{3, 2},
	{5, 0},
	{5, 3},
	{10, 0},
	{10, 5},
	{15, 10},
	{30, 0},
	{30, 20},
}

func strToTf(timeFormat string) message.StartGame {
	if timeFormat == "Custom" {
		return message.StartGame{0, 0}
	}

	for i := 0; i < len(timeFormats); i++ {
		time := fmt.Sprintf("%d", timeFormats[i].Time)
		increment := fmt.Sprintf("%d", timeFormats[i].Increment)

		if time+"+"+increment == timeFormat {
			return message.StartGame{timeFormats[i].Time, timeFormats[i].Increment}
		}
	}

	return message.StartGame{0, 0}
}

type Model struct {
	grid *grid.Model
}

func New(height, width uint) *Model {
	cfg := config.GetConfig()
	go requests.StreamIncomingEvents(cfg.Token)

	model := &Model{
		grid: grid.New(3, 4, 13, 3),
	}

	for i := 0; i < len(timeFormats); i++ {
		time := fmt.Sprintf("%d", timeFormats[i].Time)
		increment := fmt.Sprintf("%d", timeFormats[i].Increment)

		model.grid.Squares[i] = time + "+" + increment
	}

	model.grid.Squares[len(model.grid.Squares)-1] = "Custom"

	return model
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, func() tea.Msg {
				timeFormat := strToTf(m.grid.Squares[m.grid.CurrentSquare])
				return timeFormat
			}
		}
	}

	m.grid.Update(msg)

	return m, nil
}

func (m *Model) View() string {
	return m.grid.View()
}
