#pragma once
#ifndef UTILS_H_
#define UTILS_H_

#include <ncurses.h>

#define BOXIZE_WIN(win) {\
box(win, 0, 0); \
wrefresh(win); \
}

#define DESTROY_WIN(win) {\
	wclear(win); \
	wrefresh(win); \
	delwin(win); \
}

#endif
