#pragma once
#ifndef BOARD_H_
#define BOARD_H_

#include "pieces.h"
#include <ncurses.h>
#include <stdbool.h>

typedef enum { BLACK_SQUARE, WHITE_SQUARE } square_color;

typedef struct square {
  square_color color;
  piece *piece;
  WINDOW *square_window;
  unsigned col;
  unsigned row;
} square;

typedef struct board {
  WINDOW *board_window;
  square **squares;

} board;

#endif
