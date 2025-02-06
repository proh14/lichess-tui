package board

import (
	"lichess-tui/internal/tui/message"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	msg message.LoadBoard
}

var viewStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#00D8BD")).
	Width(40).
	Height(10).
	Bold(true).
	Align(lipgloss.Center).
	AlignVertical(lipgloss.Center)

func New(msg message.LoadBoard) *Model {
	return &Model{
		msg: msg,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *Model) View() string {
	str := m.msg.Data.Game.GameID + " " + m.msg.Data.Game.Opponent.Username
	return viewStyle.Render(str)
}
