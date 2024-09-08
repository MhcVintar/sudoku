package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Difficulty = uint8

const (
	Easy Difficulty = iota
	Medium
	Hard
)

type Game struct {
	eventHandler *eventHandler
	painter      *painter
	board        *board
	startTime    time.Time
}

var mistakes = 0

func NewGame(difficulty Difficulty) *Game {
	b := generateBoard(difficulty)
	c := newCursor()

	return &Game{
		eventHandler: newEventHandler(b, c),
		painter:      newPainter(b, c),
		board:        b,
		startTime:    time.Now(),
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

	if g.board.isSolved() {
		duration := time.Now().Sub(g.startTime)
		g.painter.paintStats(duration)
	}

	err = termbox.Flush()
	if err != nil {
		panic(err)
	}
}
