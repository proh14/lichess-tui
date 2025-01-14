package components

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	squareStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			AlignVertical(lipgloss.Center).
			Margin(1).
			Align(lipgloss.Center)

	foucusedSquareStyle = squareStyle.Foreground(lipgloss.Color("205"))
)

type Grid struct {
	cols          int
	rows          int
	squares       []string
	currentSquare int
	squaresWidth  int
	squaresHeight int
}

func NewGrid(rows, cols, width, height int) *Grid {
	return &Grid{
		cols:          cols,
		rows:          rows,
		squaresWidth:  width,
		squaresHeight: height,
		squares:       make([]string, rows*cols),
		currentSquare: 0,
	}
}

func (g *Grid) Init() tea.Cmd {
	return nil
}

func (g *Grid) SetSquare(row, col int, value string) {
	if row < 0 || row >= g.rows || col < 0 || col >= g.cols {
		return
	}
	g.squares[row*g.cols+col] = value
}

func (g *Grid) View() string {
	style := squareStyle.Width(g.squaresWidth).Height(g.squaresHeight)
	focusedStyle := foucusedSquareStyle.Width(g.squaresWidth).Height(g.squaresHeight)

	sb := strings.Builder{}
	tempsquares := make([]string, g.rows*g.cols)

	for square := range g.squares {
		if square == g.currentSquare {
			tempsquares[square] = focusedStyle.Render(g.squares[square])
		} else {
			tempsquares[square] = style.Render(g.squares[square])
		}
	}

	for i := 0; i < g.rows; i++ {
		sb.WriteString(
			lipgloss.JoinHorizontal(lipgloss.Center, tempsquares[i*g.cols:(i+1)*g.cols]...))
	}

	return lipgloss.JoinVertical(lipgloss.Center, sb.String())
}

func (g *Grid) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return g, nil
}
