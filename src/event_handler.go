package src

import "github.com/nsf/termbox-go"

type eventHandler struct {
	board  *board
	cursor *cursor
}

func newEventHandler(board *board, cursor *cursor) *eventHandler {
	return &eventHandler{
		board:  board,
		cursor: cursor,
	}
}

func (e *eventHandler) handle(event termbox.Event) (quit bool, err error) {
	switch event.Type {
	case termbox.EventKey:
		quit = e.handleKeyEvent(event)
	}
	return
}

func (e *eventHandler) handleKeyEvent(event termbox.Event) (quit bool) {
	switch event.Key {
	case termbox.KeyArrowUp:
		e.cursor.moveUp()
	case termbox.KeyArrowDown:
		e.cursor.moveDown()
	case termbox.KeyArrowLeft:
		e.cursor.moveLeft()
	case termbox.KeyArrowRight:
		e.cursor.moveRight()
	case termbox.KeyEsc:
		quit = true
	}

	switch event.Ch {
	case 'k':
		e.cursor.moveUp()
	case 'j':
		e.cursor.moveDown()
	case 'h':
		e.cursor.moveLeft()
	case 'l':
		e.cursor.moveRight()
	case 'q':
		quit = true
	}

	switch event.Ch {
	case '0':
		e.board.clearCell(e.cursor.boardY, e.cursor.boardX, false)
	case '1':
		e.board.setCell(e.cursor.boardY, e.cursor.boardX, 1, false)
	case '2':
		e.board.setCell(e.cursor.boardY, e.cursor.boardX, 2, false)
	case '3':
		e.board.setCell(e.cursor.boardY, e.cursor.boardX, 3, false)
	case '4':
		e.board.setCell(e.cursor.boardY, e.cursor.boardX, 4, false)
	case '5':
		e.board.setCell(e.cursor.boardY, e.cursor.boardX, 5, false)
	case '6':
		e.board.setCell(e.cursor.boardY, e.cursor.boardX, 6, false)
	case '7':
		e.board.setCell(e.cursor.boardY, e.cursor.boardX, 7, false)
	case '8':
		e.board.setCell(e.cursor.boardY, e.cursor.boardX, 8, false)
	case '9':
		e.board.setCell(e.cursor.boardY, e.cursor.boardX, 9, false)
	}

	return
}
