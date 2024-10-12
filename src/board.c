#include <board.h>
#include <utils.h>
#include <stdlib.h>

struct board board;

void init_board() {
	board.board_window = create_newwin(0, 0, 16, 8);

	board.squares = malloc(8 * sizeof(struct piece[8]));
	for (size_t i = 0; i < 8; ++i) {
		board.squares[i] = malloc(8 * sizeof(struct piece));
	}
}
