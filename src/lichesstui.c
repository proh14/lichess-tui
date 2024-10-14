#include <ncurses.h>
#include <board.h>

int main(void) {
	initscr();
	
	noecho();
	cbreak(); 

	refresh();

	init_board();

    getchar();

	endwin();

	return 0;
}
