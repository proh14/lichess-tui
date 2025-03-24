package grid

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	WHITE = "white"
	BLACK = "black"
)

var (
	squareOddStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			AlignVertical(lipgloss.Center).
			Align(lipgloss.Center)
	squareEvenStyle = squareOddStyle.Bold(true).
			BorderForeground(lipgloss.Color("#00D8BD")).
			Foreground(lipgloss.Color("#00D8BD"))

	foucusedSquareStyle = squareOddStyle.Bold(true).
				BorderForeground(lipgloss.Color("#a834eb")).
				Foreground(lipgloss.Color("#a834eb"))
)

type Model struct {
	Cols          int
	Rows          int
	Squares       []string
	CurrentSquare int
	SquaresWidth  int
	SquaresHeight int
	Color         string
}

func New(cols, rows, width, height int, color string) *Model {
	return &Model{
		Cols:          cols,
		Rows:          rows,
		SquaresWidth:  width,
		SquaresHeight: height,
		Squares:       make([]string, rows*cols),
		CurrentSquare: 0,
		Color:         color,
	}
}

func (g *Model) Init() tea.Cmd {
	return nil
}

func IsWhite(color string) int {
	if color == WHITE {
		return 1
	} else {
		return 0
	}
}

func (g *Model) View() string {
	styleOdd := squareOddStyle.Width(g.SquaresWidth).Height(g.SquaresHeight)
	styleEven := squareEvenStyle.Width(g.SquaresWidth).Height(g.SquaresHeight)
	focusedStyle := foucusedSquareStyle.Width(g.SquaresWidth).Height(g.SquaresHeight)

	tempsquares := make([]string, g.Rows*g.Cols)
	for i, square := range g.Squares {
		n := i + 1
		row := (n - 1) / g.Rows
		col := (n - 1) % g.Rows
		if i == g.CurrentSquare {
			tempsquares[i] = focusedStyle.Render(square)
		} else if (row+col)%2 == IsWhite(g.Color) {
			tempsquares[i] = styleEven.Render(square)
		} else {
			tempsquares[i] = styleOdd.Render(square)
		}
	}

	rows := make([]string, g.Rows)

	for i := range rows {
		rows[i] = lipgloss.JoinHorizontal(lipgloss.Center, tempsquares[i*g.Cols:(i+1)*g.Cols]...)
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (g *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			g.Up()
		case "down":
			g.Down()
		case "left":
			g.Left()
		case "right":
			g.Right()
		}
	}
	return g, nil
}

func (g *Model) Left() {
	if g.CurrentSquare > 0 {
		g.CurrentSquare--
	}
}

func (g *Model) Right() {
	if g.CurrentSquare < g.Rows*g.Cols-1 {
		g.CurrentSquare++
	}
}

func (g *Model) Up() {
	if g.CurrentSquare-g.Cols >= 0 {
		g.CurrentSquare -= g.Cols
	}
}

func (g *Model) Down() {
	if g.CurrentSquare+g.Cols < g.Rows*g.Cols {
		g.CurrentSquare += g.Cols
	}
}
