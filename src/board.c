#include <board.h>
#include <utils.h>
#include <stdlib.h>

struct board board;

void init_board() {
	board.board_window = newwin(SQUARE_SIZE * 8, SQUARE_SIZE * 8, 0, 0);
	BOXIZE_WIN(board.board_window);

	board.squares = malloc(8 * sizeof(struct piece[8]));

	for (short int i = 0; i < 8; ++i) {
		board.squares[i] = malloc(8 * sizeof(struct piece));

		for (short int j = 0; j < 8; ++j) {
			square *square = &board.squares[i][j];
		
			square->square_window = derwin(board.board_window, SQUARE_SIZE, SQUARE_SIZE, i * SQUARE_SIZE, j * SQUARE_SIZE);
			BOXIZE_WIN(square->square_window);
		}
	}
}
