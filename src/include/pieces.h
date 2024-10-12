#pragma once
#ifndef PIECES_H_
#define PIECES_H_

#include <stdbool.h>

typedef enum {
  ROOK,
  BISHOP,
  PAWN,
  KING,
  QUEEN,
} piece_type;

typedef enum { BLACK, WHITE } piece_color;

typedef struct piece {
  piece_type type;
  piece_color color;
  bool has_moved;
} piece;

#endif
