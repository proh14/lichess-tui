package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/proh14/lichess-tui/internal/tui/components"
)

type viewState uint

const (
	FormView viewState = iota
	QuickGameView
	GameView
)

type Model struct {
	viewState      viewState
	quickGameModel *components.QuickGame
}

func NewModel() *Model {
	return &Model{
		viewState:      QuickGameView,
		quickGameModel: components.NewQuickGameModel(),
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
	return m.quickGameModel.View()
}
