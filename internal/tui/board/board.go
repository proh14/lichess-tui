package board

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	time      uint
	increment uint
	gameID    string
}

var viewStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#00D8BD")).
	Width(40).
	Height(10).
	Bold(true).
	Align(lipgloss.Center).
	AlignVertical(lipgloss.Center)

func New(time, increment uint, gameID string) *Model {
	return &Model{
		time:      time,
		increment: increment,
		gameID:    gameID,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *Model) View() string {
	str := fmt.Sprintf(m.gameID)
	return viewStyle.Render(str)
}
