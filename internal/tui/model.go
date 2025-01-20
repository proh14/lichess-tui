package tui

import (
	"strings"

	"lichess-tui/internal/config"
	"lichess-tui/internal/requests"
	"lichess-tui/internal/tui/message"
	"lichess-tui/internal/tui/quickgame"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type viewState uint

const (
	QuickGameView = iota
	GameView
)

type Model struct {
	viewState      viewState
	quickGameModel *quickgame.Model
	title          string
	profile        requests.Profile
	status         string
	loaded         bool
	height         int
	width          int
}

func NewModel() *Model {
	cfg := config.GetConfig()
	profile := requests.GetProfile(cfg.Token)

	return &Model{
		viewState:      QuickGameView,
		quickGameModel: quickgame.New(0, 0),
		loaded:         false,
		profile:        profile,
		status:         "\nPress 'q' to quit",
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.SetWindowTitle("Lichess TUI")
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.loaded = true
	case message.StartGame:
		cfg := config.GetConfig()
		go requests.SeekGame(requests.SeekGameConfig{Time: 10, Increment: 0}, cfg.Token)
		m.status = "\nStarting the game..."
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	_, cmd := m.quickGameModel.Update(msg)

	return m, cmd
}

func (m *Model) View() string {
	if !m.loaded {
		return "Loading..."
	}

	sb := strings.Builder{}

	m.title = m.profile.ID + "\n"
	m.title += strings.Repeat("─", m.width-1)
	m.title += "\n"

	sb.WriteString(m.title)

	sb.WriteString(lipgloss.Place(
		m.width,
		m.height-4,
		lipgloss.Center,
		lipgloss.Center,
		m.quickGameModel.View(),
	),
	)

	sb.WriteString(m.status)

	return sb.String()
}
