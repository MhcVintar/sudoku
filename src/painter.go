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
	boardLines := []string{
		"┌─────────┬─────────┬─────────┐",
		"│         │         │         │",
		"│         │         │         │",
		"│         │         │         │",
		"├─────────┼─────────┼─────────┤",
		"│         │         │         │",
		"│         │         │         │",
		"│         │         │         │",
		"├─────────┼─────────┼─────────┤",
		"│         │         │         │",
		"│         │         │         │",
		"│         │         │         │",
		"└─────────┴─────────┴─────────┘",
	}

	for y, line := range boardLines {
		x := 0
		for _, character := range line {
			termbox.SetCell(x, y, character, termbox.ColorDefault, termbox.ColorDefault)
			x++
		}
	}

	x, y := 2, 1
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			value := p.board.getCell(row, column).value
			isGenerated := p.board.getCell(row, column).isGenerated
			isValid := p.board.getCell(row, column).isValid

			var character rune
			if value == 0 {
				character = ' '
			} else {
				character = rune(value + '0')
			}

			var digitStyle termbox.Attribute
			if isGenerated {
				digitStyle = termbox.ColorLightBlue | termbox.AttrBold
			} else if !isValid {
				digitStyle = termbox.ColorLightRed | termbox.AttrCursive
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

	if p.board.isSolved() {
		p.paintVictory()
	}
}

func (p *painter) paintVictory() {
	const offsetX = 8
	const offsetY = 5

	lines := []string{
		"┌─┴─────────┴─┐",
		"│   Victory   │",
		"└─┬─────────┬─┘",
	}

	for y, line := range lines {
		x := 0
		for _, character := range line {
			termbox.SetCell(x+offsetX, y+offsetY, character, termbox.ColorLightGreen, termbox.ColorDefault)
			x++
		}
	}
}
