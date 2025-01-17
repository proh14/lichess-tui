package quickgame

import (
	"github.com/proh14/lichess-tui/internal/tui/grid"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	grid *grid.Model
}

func New(height, width uint) *Model {
	return &Model{
		grid: grid.New(3, 3, 11, 4),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *Model) View() string {
	return m.grid.View()
}
