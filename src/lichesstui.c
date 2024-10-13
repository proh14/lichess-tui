#include <ncurses.h>
#include <board.h>

int main(void) {
	initscr();
	cbreak(); 

	refresh();

	init_board();

	while (true) {
	}

	/*printf("%d", board.board_window->width);*/

	endwin();

	return 0;
}
