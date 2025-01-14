package components

import (
	tea "github.com/charmbracelet/bubbletea"
)

type QuickGame struct {
	grid *Grid
}

func NewQuickGameModel() *QuickGame {
	return &QuickGame{
		grid: NewGrid(3, 3, 7, 3),
	}
}

func (m QuickGame) Init() tea.Cmd {
	return nil
}

func (m QuickGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m QuickGame) View() string {
	return m.grid.View()
}
