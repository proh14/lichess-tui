package quickgame

import (
	"fmt"

	"lichess-tui/internal/tui/grid"
	"lichess-tui/internal/tui/message"

	tea "github.com/charmbracelet/bubbletea"
)

type timeFormat struct {
	time      uint
	increment uint
}

var timeFormats = [...]timeFormat{
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

func strToTf(timeFormat string) (time uint, increment uint) {
	if timeFormat == "Custom" {
		return 0, 0
	}

	for i := 0; i < len(timeFormats); i++ {
		time := fmt.Sprintf("%d", timeFormats[i].time)
		increment := fmt.Sprintf("%d", timeFormats[i].increment)

		if time+"+"+increment == timeFormat {
			return timeFormats[i].time, timeFormats[i].increment
		}
	}

	return 0, 0
}

type Model struct {
	grid *grid.Model
}

func New(height, width uint) *Model {
	model := &Model{
		grid: grid.New(3, 4, 13, 3),
	}

	for i := 0; i < len(timeFormats); i++ {
		time := fmt.Sprintf("%d", timeFormats[i].time)
		increment := fmt.Sprintf("%d", timeFormats[i].increment)

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
				time, increment := strToTf(m.grid.Squares[m.grid.CurrentSquare])
				return message.StartGame{Time: time, Increment: increment}
			}
		}
	}

	m.grid.Update(msg)

	return m, nil
}

func (m *Model) View() string {
	return m.grid.View()
}
