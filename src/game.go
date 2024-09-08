package src

import "github.com/nsf/termbox-go"

type Difficulty = uint8

const (
	Easy Difficulty = iota
	Medium
	Hard
)

type Game struct {
	eventHandler *eventHandler
	painter      *painter
}

func NewGame(difficulty Difficulty) *Game {
	b := generateBoard(difficulty)
	c := newCursor()

	return &Game{
		eventHandler: newEventHandler(b, c),
		painter:      newPainter(b, c),
	}
}

func (g *Game) Run() {
	for {
		g.repaint()

		quit, err := g.eventHandler.handle(termbox.PollEvent())
		if err != nil {
			panic(err)
		}
		if quit {
			return
		}
	}
}

func (g *Game) repaint() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}

	g.painter.paintBoard()
	g.painter.paintCursor()
	g.painter.paintCommands()

	err = termbox.Flush()
	if err != nil {
		panic(err)
	}
}
