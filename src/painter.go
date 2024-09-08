package src

import (
	"github.com/nsf/termbox-go"
)

type painter struct {
	board  *board
	cursor *cursor
}

func newPainter(board *board, cursor *cursor) *painter {
	return &painter{
		board:  board,
		cursor: cursor,
	}
}

func (p *painter) paintCursor() {
	currentCell := termbox.GetCell(p.cursor.screenX, p.cursor.screenY)
	termbox.SetCell(p.cursor.screenX, p.cursor.screenY, currentCell.Ch, termbox.ColorBlack, termbox.ColorWhite)
}

func (p *painter) paintBoard() {
	// Draw top bar
	termbox.SetCell(0, 0, '┌', termbox.ColorDefault, termbox.ColorDefault)
	for i := 1; i <= 29; i++ {
		if i == 10 || i == 20 {
			termbox.SetCell(i, 0, '┬', termbox.ColorDefault, termbox.ColorDefault)
		} else {
			termbox.SetCell(i, 0, '─', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	termbox.SetCell(30, 0, '┐', termbox.ColorDefault, termbox.ColorDefault)

	// Draw bottom bar
	termbox.SetCell(0, 12, '└', termbox.ColorDefault, termbox.ColorDefault)
	for i := 1; i <= 29; i++ {
		if i == 10 || i == 20 {
			termbox.SetCell(i, 12, '┴', termbox.ColorDefault, termbox.ColorDefault)
		} else {
			termbox.SetCell(i, 12, '─', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	termbox.SetCell(30, 12, '┘', termbox.ColorDefault, termbox.ColorDefault)

	// Draw left bar
	for i := 1; i <= 11; i++ {
		if i == 4 || i == 8 {
			termbox.SetCell(0, i, '├', termbox.ColorDefault, termbox.ColorDefault)
		} else {
			termbox.SetCell(0, i, '│', termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	// Draw right bar
	for i := 1; i <= 11; i++ {
		if i == 4 || i == 8 {
			termbox.SetCell(30, i, '┤', termbox.ColorDefault, termbox.ColorDefault)
		} else {
			termbox.SetCell(30, i, '│', termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	// Draw middle 2 horizontal bars
	for i := 1; i <= 29; i++ {
		if i == 10 || i == 20 {
			termbox.SetCell(i, 4, '┼', termbox.ColorDefault, termbox.ColorDefault)
			termbox.SetCell(i, 8, '┼', termbox.ColorDefault, termbox.ColorDefault)
		} else {
			termbox.SetCell(i, 4, '─', termbox.ColorDefault, termbox.ColorDefault)
			termbox.SetCell(i, 8, '─', termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	// Draw middle 2 vertical bars
	for i := 1; i <= 11; i++ {
		if i != 4 && i != 8 {
			termbox.SetCell(10, i, '│', termbox.ColorDefault, termbox.ColorDefault)
			termbox.SetCell(20, i, '│', termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	// Draw digits
	x, y := 2, 1
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			value := p.board.getValue(row, column)
			isGenerated := p.board.isGenerated(row, column)

			var character rune
			if value == 0 {
				character = ' '
			} else {
				character = rune(value + '0')
			}

			var digitStyle termbox.Attribute
			if isGenerated {
				digitStyle = termbox.ColorLightBlue | termbox.AttrBold
			} else {
				digitStyle = termbox.ColorDefault
			}

			termbox.SetCell(x, y, character, digitStyle, termbox.ColorDefault)

			if column == 2 || column == 5 {
				x += 4
			} else {
				x += 3
			}
		}

		x = 2
		if row == 2 || row == 5 {
			y += 2
		} else {
			y++
		}
	}

	// Draw victory message
	if p.board.isSolved() {
		// Draw top bar
		termbox.SetCell(8, 5, '┌', termbox.ColorLightGreen, termbox.ColorDefault)
		for i := 9; i <= 21; i++ {
			if i == 10 || i == 20 {
				termbox.SetCell(i, 5, '┴', termbox.ColorLightGreen, termbox.ColorDefault)
			} else {
				termbox.SetCell(i, 5, '─', termbox.ColorLightGreen, termbox.ColorDefault)
			}
		}
		termbox.SetCell(22, 5, '┐', termbox.ColorLightGreen, termbox.ColorDefault)

		// Draw bottom bar
		termbox.SetCell(8, 7, '└', termbox.ColorLightGreen, termbox.ColorDefault)
		for i := 9; i <= 21; i++ {
			if i == 10 || i == 20 {
				termbox.SetCell(i, 7, '┬', termbox.ColorLightGreen, termbox.ColorDefault)
			} else {
				termbox.SetCell(i, 7, '─', termbox.ColorLightGreen, termbox.ColorDefault)
			}
		}
		termbox.SetCell(22, 7, '┘', termbox.ColorLightGreen, termbox.ColorDefault)

		// Draw left bar
		termbox.SetCell(8, 6, '│', termbox.ColorLightGreen, termbox.ColorDefault)

		// Draw right bar
		termbox.SetCell(22, 6, '│', termbox.ColorLightGreen, termbox.ColorDefault)

		// Draw text
		termbox.SetCell(9, 6, ' ', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(10, 6, ' ', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(11, 6, ' ', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(12, 6, 'V', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(13, 6, 'i', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(14, 6, 'c', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(15, 6, 't', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(16, 6, 'o', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(17, 6, 'r', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(18, 6, 'y', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(19, 6, ' ', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(20, 6, ' ', termbox.ColorLightGreen, termbox.ColorDefault)
		termbox.SetCell(21, 6, ' ', termbox.ColorLightGreen, termbox.ColorDefault)
	}
}
