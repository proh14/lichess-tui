#include <ncurses.h>

int main(void) {
  initscr();

  printw("HELLO C!!!!!!!!!!!!!");
  refresh();
  getch();
  endwin();
}
