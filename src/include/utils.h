#pragma once
#ifndef UTILS_H_
#define UTILS_H_

#include <ncurses.h>

WINDOW *create_newwin(int x, int y, int width, int height);

void destroy_win(WINDOW *local_win);

#endif
