package board

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/notnil/chess"
)

var (
	whiteSquareStyle = lipgloss.NewStyle().
				Width(6).
				Height(3).
				Align(lipgloss.Center).
				Background(lipgloss.Color("#f0f0f0")).
				Foreground(lipgloss.Color("#000000")).
				Bold(true)
	blackSquareStyle = lipgloss.NewStyle().
				Align(lipgloss.Center).
				Width(6).
				Height(3).
				Background(lipgloss.Color("#333333")).
				Bold(true)
)

type Model struct {
	Board *chess.Board
}

func New() *Model {
	b := chess.NewBoard(chess.StartingPosition().Board().SquareMap())
	return &Model{
		Board: b,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *Model) View() string {
	squares := make([]string, 8*8)

	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			s := ""
			p := m.Board.Piece(chess.NewSquare(chess.File(f), chess.Rank(r)))
			if p != chess.NoPiece {
				s = p.String()
			}
			s += " "
			if (r+f)%2 == 0 {
				squares[r*8+f] = whiteSquareStyle.Render(s)
			} else {
				squares[r*8+f] = blackSquareStyle.Render(s)
			}
		}
	}

	rows := make([]string, 8)

	for r := 0; r < 8; r++ {
		rows[r] = lipgloss.JoinHorizontal(lipgloss.Center, squares[r*8:(r+1)*8]...)
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
