package starting

import (
	"fmt"
	"lichess-tui/internal/config"
	"lichess-tui/internal/requests"
	"lichess-tui/internal/tui/message"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	time      uint
	increment uint
}

var viewStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#00D8BD")).
	Width(40).
	Height(10).
	Bold(true).
	Align(lipgloss.Center).
	AlignVertical(lipgloss.Center)

func New(time, increment uint) *Model {
	cfg := config.GetConfig()
	go requests.StreamIncomingEvents(cfg.Token)

	return &Model{
		time:      time,
		increment: increment,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	fmt.Println(requests.IncomingEventsData)
	if requests.IncomingEventsData.Type != "gameStart" {
		return m, nil
	}

	return m, func() tea.Msg {
		return message.LoadBoard{
			Time:      m.time,
			Increment: m.increment,
		}
	}
}

func (m *Model) View() string {
	str := fmt.Sprintf("Matchmaking...\n[%d+%d]", m.time, m.increment)
	return viewStyle.Render(str)
}
