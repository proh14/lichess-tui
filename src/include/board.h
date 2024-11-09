#pragma once
#ifndef BOARD_H_
#define BOARD_H_

#include <pieces.h>
#include <ncurses.h>
#include <stdbool.h>

#define SQUARE_WIDTH_RATIO 2
#define SQUARE_HEIGHT 3
#define SQUARE_WIDTH SQUARE_WIDTH_RATIO * SQUARE_HEIGHT

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
