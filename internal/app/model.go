package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/proh14/lichess-tui/internal/config"
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
  mycfg := config.GetConfig()
  token := mycfg.Token
	return fmt.Sprintf("Token is %s", token)
}
