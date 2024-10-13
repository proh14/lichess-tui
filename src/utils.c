#include <ncurses.h>

WINDOW *create_newwin(int x, int y, int width, int height)
{	WINDOW *local_win;

	local_win = newwin(height, width, y, x);
	box(local_win, 0, 0);
	wrefresh(local_win);

	return local_win;
}

void destroy_win(WINDOW *local_win)
{	
	wborder(local_win, ' ', ' ', ' ',' ',' ',' ',' ',' ');
	wrefresh(local_win);
	delwin(local_win);
}
