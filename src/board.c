#include <board.h>
#include <utils.h>
#include <stdlib.h>

struct board board;

void init_board() {
	board.squares = malloc(8 * sizeof(struct piece[8]));

	for (size_t i = 0; i < 8; ++i) {
		board.squares[i] = malloc(8 * sizeof(struct piece));

		for (size_t j = 0; j < 8; ++j) {
			square *square = &board.squares[i][j];
			
			square->square_window = create_newwin(i * SQUARE_SIZE, j * SQUARE_SIZE, SQUARE_SIZE, SQUARE_SIZE);
		}
	}
}
