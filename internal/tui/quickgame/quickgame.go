package quickgame

import (
	"github.com/proh14/lichess-tui/internal/tui/grid"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	grid *grid.Model
}

func New(height, width uint) *Model {
	model := &Model{
		grid: grid.New(3, 4, 13, 3),
	}

	timeFormats := [...]string{
		"1+0",
		"2+1",
		"3+0",
		"3+2",
		"5+0",
		"5+3",
		"10+0",
		"10+5",
		"15+10",
		"30+0",
		"30+20",
		"Custom.",
	}

	for i := 0; i < 12; i++ {
		model.grid.Squares[i] = timeFormats[i]
	}

	return model
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.grid.Update(msg)
	return m, nil
}

func (m *Model) View() string {
	return m.grid.View()
}
