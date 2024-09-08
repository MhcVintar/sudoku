package main

type cursor struct {
	screenX, screenY, boardX, boardY int
}

func newCursor() *cursor {
	return &cursor{
		screenX: 2,
		screenY: 1,
		boardX:  0,
		boardY:  0,
	}
}

func (c *cursor) moveLeft() {
	if c.screenX == 2 {
		return
	}
	c.screenX -= 4
	c.boardX--
}

func (c *cursor) moveRight() {
	if c.screenX == 34 {
		return
	}
	c.screenX += 4
	c.boardX++
}

func (c *cursor) moveUp() {
	if c.screenY == 1 {
		return
	}
	c.screenY -= 2
	c.boardY--
}

func (c *cursor) moveDown() {
	if c.screenY == 17 {
		return
	}
	c.screenY += 2
	c.boardY++
}
