#pragma once
#ifndef BOARD_H_
#define BOARD_H_

#include <pieces.h>
#include <ncurses.h>
#include <stdbool.h>

#define SQUARE_SIZE 3 // The size is a unit in charge of both the width (x) and the height (y): x = y

typedef enum { BLACK_SQUARE, WHITE_SQUARE } square_color;

typedef struct square {
  square_color color;
  piece *piece;
  WINDOW *square_window;
  unsigned col;
  unsigned row;
} square;

struct board {
  square **squares;
  WINDOW *board_window;
};

void init_board();

#endif
