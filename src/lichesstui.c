#include <ncurses.h>
#include <board.h>

int main(void) {
	initscr();
	cbreak(); 

	refresh();

	init_board();

	while (true) {

	}

	endwin();

	return 0;
}
