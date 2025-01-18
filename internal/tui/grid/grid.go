package grid

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	squareStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			AlignVertical(lipgloss.Center).
			Align(lipgloss.Center)

	foucusedSquareStyle = squareStyle.Foreground(lipgloss.Color("205"))
)

type Model struct {
	Cols          int
	Rows          int
	Squares       []string
	CurrentSquare int
	SquaresWidth  int
	SquaresHeight int
}

func New(cols, rows, width, height int) *Model {
	return &Model{
		Cols:          cols,
		Rows:          rows,
		SquaresWidth:  width,
		SquaresHeight: height,
		Squares:       make([]string, rows*cols),
		CurrentSquare: 0,
	}
}

func (g *Model) Init() tea.Cmd {
	return nil
}

func (g *Model) View() string {
	style := squareStyle.Width(g.SquaresWidth).Height(g.SquaresHeight)
	focusedStyle := foucusedSquareStyle.Width(g.SquaresWidth).Height(g.SquaresHeight)

	tempsquares := make([]string, len(g.Squares))
	for square := range g.Squares {
		if square == g.CurrentSquare {
			tempsquares[square] = focusedStyle.Render(g.Squares[square])
		} else {
			tempsquares[square] = style.Render(g.Squares[square])
		}
	}

	rows := make([]string, g.Rows)
	for i := 0; i < g.Rows; i++ {
		rows[i] = lipgloss.JoinHorizontal(lipgloss.Center, tempsquares[i*g.Cols:(i+1)*g.Cols]...)
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (g *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return g, nil
}

func (g *Model) Left() {
	if g.CurrentSquare > 0 {
		g.CurrentSquare--
	}
}

func (g *Model) Right() {
	if g.CurrentSquare < len(g.Squares)-1 {
		g.CurrentSquare++
	}
}

func (g *Model) Up() {
	if g.CurrentSquare-g.Cols >= 0 {
		g.CurrentSquare -= g.Cols
	}
}

func (g *Model) Down() {
	if g.CurrentSquare+g.Cols < len(g.Squares) {
		g.CurrentSquare += g.Cols
	}
}
