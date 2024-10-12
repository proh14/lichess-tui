#include "board.h"
#include <stdlib.h>

struct board board;

void init_board() {
	board.board_window = 
		newwin(5, 5, 5, 5);
	board.squares = malloc(2 * sizeof(int));
}
