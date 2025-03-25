package board

import (
	"lichess-tui/internal/tui/grid"
	"lichess-tui/internal/tui/message"
	"math"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/notnil/chess"
)

const (
	ChoosingStarting int = iota
	ChoosingEnding
)

type Model struct {
	game *chess.Game
	grid *grid.Model
	startingSquare chess.Square
	endingSquare chess.Square
	mode int
	msg  message.LoadBoard
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
	fen, _ := chess.FEN(msg.Data.Game.Fen)
	game := chess.NewGame(fen)

	board := game.Position().Board()
	// msg.data.game.color
	color := msg.Data.Game.Color
	model := &Model{
		game: game,
		grid: grid.New(8, 8, 5, 2, color),
		msg:  msg,
	}

	switch color {
	case grid.BLACK:
		board = board.Flip(chess.LeftRight)
	case grid.WHITE:
		board = board.Flip(chess.UpDown)
	}

	for i := range 8 {
		for j := range 8 {
			index := chess.NewSquare(chess.File(j), chess.Rank(i))
			string := ""
			switch board.Piece(index) {
			case chess.NoPiece:
				string = ""
			case chess.BlackKing:
				string = "♔"
			case chess.BlackQueen:
				string = "♕"
			case chess.BlackRook:
				string = "♖"
			case chess.BlackBishop:
				string = "♗"
			case chess.BlackKnight:
				string = "♘"
			case chess.BlackPawn:
				string = "♙"
			case chess.WhiteKing:
				string = "♚"
			case chess.WhiteQueen:
				string = "♛"
			case chess.WhiteRook:
				string = "♜"
			case chess.WhiteBishop:
				string = "♝"
			case chess.WhiteKnight:
				string = "♞"
			case chess.WhitePawn:
				string = "♟"
			}

			model.grid.Squares[index] = string
		}
	}

	return model
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.grid.Update(msg)
	
	x := math.Abs(7 - math.Floor(float64(m.grid.CurrentSquare) / 8.0));
	y := m.grid.CurrentSquare % 8;

	index := chess.NewSquare(chess.File(y), chess.Rank(x));

	enter_pressed := false;
	if msg, ok := msg.(tea.KeyMsg); ok {
		if msg.Type == tea.KeyEnter {
			enter_pressed = true;
		}
	}

	switch m.mode {
		case ChoosingStarting:
			m.startingSquare = index;
			if enter_pressed {
				m.mode = ChoosingEnding;
			}
		case ChoosingEnding:
			m.endingSquare = index;
			if enter_pressed {
				move, _ := chess.LongAlgebraicNotation{}.Decode(m.game.Position(), m.startingSquare.String() + m.endingSquare.String());
				m.game.Move(move);
				fmt.Println(m.game.Position());

				m.mode = ChoosingStarting;
			}
	}

	return m, nil
}

func (m *Model) View() string {
	return m.grid.View()
}
