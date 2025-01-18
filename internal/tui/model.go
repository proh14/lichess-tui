package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/proh14/lichess-tui/internal/config"
	"github.com/proh14/lichess-tui/internal/requests"
	"github.com/proh14/lichess-tui/internal/tui/quickgame"
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
	status         string
	loaded         bool
	height         int
	width          int
}

func NewModel() *Model {
	return &Model{
		viewState:      QuickGameView,
		quickGameModel: quickgame.New(0, 0),
		loaded:         false,
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
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	m.quickGameModel.Update(msg)

	return m, nil
}

func (m *Model) View() string {
	if !m.loaded {
		return "Loading..."
	}

	sb := strings.Builder{}

	cfg := config.GetConfig()
	data := requests.GetProfile(cfg.Token)

	m.title = data.ID + "\n"
	m.title += strings.Repeat("─", m.width-1)
	m.title += "\n"

	m.status = "\n"
	m.status += strings.Repeat("─", m.width-1)
	m.status += "\nHello world :)"

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
