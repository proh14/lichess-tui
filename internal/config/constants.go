package config

const (
	NAME      string = "Lichess TUI"
	CELL_SIZE int    = 1
)

type Piece uint8

const (
	Pawn   Piece = 0
	Rook   Piece = 1
	Knight Piece = 2
	Bishop Piece = 3
	Queen  Piece = 4
	King   Piece = 5
)

const (
	Pawn_Str   string = "      ███      \n    ███████    \n      ███      \n    ███████    \n  ███████████  "
	Knight_Str string = "     ██████    \n  ██████████   \n     █████     \n    ███████    \n  ███████████  "
	Bishop_Str string = "    ██ ████    \n     ██ ██     \n      ███      \n    ███████    \n  ███████████  "
	RookStr    string = "  ██  ███  ██  \n  ███████████  \n     █████     \n   █████████   \n  ███████████  "
	QueenStr   string = " ██  ██ ██  ██ \n   █████████   \n     █████     \n    ███████    \n  ███████████  "
	KingStr    string = "     █████     \n      ███      \n  ███████████  \n    ███████    \n  ███████████  "
)
