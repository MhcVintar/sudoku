package src

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
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

	var digitStyle termbox.Attribute
	if p.board.getCell(p.cursor.boardY, p.cursor.boardX).isGenerated {
		digitStyle = termbox.ColorBlue | termbox.AttrBold
	} else if !p.board.getCell(p.cursor.boardY, p.cursor.boardX).isValid {
		digitStyle = termbox.ColorLightRed | termbox.AttrBold
	} else {
		digitStyle = termbox.ColorBlack
	}

	termbox.SetCell(p.cursor.screenX-1, p.cursor.screenY, ' ', digitStyle, termbox.ColorWhite)
	termbox.SetCell(p.cursor.screenX, p.cursor.screenY, currentCell.Ch, digitStyle, termbox.ColorWhite)
	termbox.SetCell(p.cursor.screenX+1, p.cursor.screenY, ' ', digitStyle, termbox.ColorWhite)
}

func (p *painter) paintBoard() {
	boardBoldLines := []string{
		"┌───────────┬───────────┬───────────┐",
		"│           │           │           │",
		"│           │           │           │",
		"│           │           │           │",
		"│           │           │           │",
		"│           │           │           │",
		"├───────────┼───────────┼───────────┤",
		"│           │           │           │",
		"│           │           │           │",
		"│           │           │           │",
		"│           │           │           │",
		"│           │           │           │",
		"├───────────┼───────────┼───────────┤",
		"│           │           │           │",
		"│           │           │           │",
		"│           │           │           │",
		"│           │           │           │",
		"│           │           │           │",
		"└───────────┴───────────┴───────────┘",
	}
	for y, line := range boardBoldLines {
		x := 0
		for _, character := range line {
			termbox.SetCell(x, y, character, termbox.ColorLightBlue|termbox.AttrBold, termbox.ColorDefault)
			x++
		}
	}

	boardRegularLines := []string{
		".....................................",
		".   │   │   .   │   │   .   │   │   .",
		".───┼───┼───.───│───│───.───│───│───.",
		".   │   │   .   │   │   .   │   │   .",
		".───┼───┼───.───┼───┼───.───┼───┼───.",
		".   │   │   .   │   │   .   │   │   .",
		".....................................",
		".   │   │   .   │   │   .   │   │   .",
		".───┼───┼───.───┼───┼───.───┼───┼───.",
		".   │   │   .   │   │   .   │   │   .",
		".───┼───┼───.───┼───┼───.───┼───┼───.",
		".   │   │   .   │   │   .   │   │   .",
		".....................................",
		".   │   │   .   │   │   .   │   │   .",
		".───┼───┼───.───┼───┼───.───┼───┼───.",
		".   │   │   .   │   │   .   │   │   .",
		".───┼───┼───.───┼───┼───.───┼───┼───.",
		".   │   │   .   │   │   .   │   │   .",
		".....................................",
	}

	for y, line := range boardRegularLines {
		x := 0
		for _, character := range line {
			if character != '.' {
				termbox.SetCell(x, y, character, termbox.ColorLightBlue, termbox.ColorDefault)
			}
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
				digitStyle = termbox.ColorLightRed | termbox.AttrBold
			} else {
				digitStyle = termbox.ColorDefault
			}

			termbox.SetCell(x, y, character, digitStyle, termbox.ColorDefault)

			x += 4
		}
		x = 2
		y += 2
	}
}

func (p *painter) paintStats(duration time.Duration) {
	const offsetX = 0
	const offsetY = 20

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	lines := []string{
		"┌────────────────────────┐",
		"│         Victory        │",
		fmt.Sprintf("│     Time: %02d:%02d:%02d     │", hours, minutes, seconds),
		fmt.Sprintf("│     Mistakes: %-*d      │", 3, mistakes),
		"└────────────────────────┘",
	}

	for y, line := range lines {
		x := 0
		for _, character := range line {
			termbox.SetCell(x+offsetX, y+offsetY, character, termbox.ColorLightGreen|termbox.AttrBold, termbox.ColorDefault)
			x++
		}
	}
}

func (p *painter) paintCommands() {
	const offsetX = 40
	const offsetY = 1

	lines := []string{
		"Arrows: move      ",
		"1-9:    set cell  ",
		"0:      clear cell",
		"q:      quit      ",
	}

	for y, line := range lines {
		x := 0
		for _, character := range line {
			termbox.SetCell(x+offsetX, y+offsetY, character, termbox.ColorDefault|termbox.AttrCursive, termbox.ColorDefault)
			x++
		}
	}
}
