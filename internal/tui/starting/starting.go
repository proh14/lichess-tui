package starting

import (
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	time      uint
	increment uint
	timer     timer.Model
}

func (m *Model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *Model) View() string {
	return
}
