package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

type viewState uint

const (
	FormView viewState = iota
	QuickGameView
	GameView
)

type Model struct {
	viewState viewState
}

func NewModel() *Model {
	return &Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	return "Hello, World!"
}
