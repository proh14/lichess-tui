#include <board.h>
#include <utils.h>
#include <stdlib.h>

struct board board;

void init_board() {
	short int board_x;
	short int board_y;
	getmaxyx(stdscr, board_y, board_x);

	board.board_window = newwin(SQUARE_HEIGHT * 8, SQUARE_WIDTH * 8, board_y / 2 - SQUARE_HEIGHT * 4, board_x / 2 - SQUARE_WIDTH * 4);
	BOXIZE_WIN(board.board_window);

	board.squares = malloc(8 * sizeof(struct piece[8]));

	for (short int i = 0; i < 8; ++i) {
		board.squares[i] = malloc(8 * sizeof(struct piece));

		for (short int j = 0; j < 8; ++j) {
			square *square = &board.squares[i][j];
		
			square->square_window = derwin(board.board_window, SQUARE_HEIGHT, SQUARE_WIDTH, i * SQUARE_HEIGHT, j * SQUARE_WIDTH);
			BOXIZE_WIN(square->square_window);
		}
	}
}
