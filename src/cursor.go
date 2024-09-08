package src

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
	switch {
	case c.screenX == 2:
		return
	case c.screenX == 12 || c.screenX == 22:
		c.screenX -= 4
	default:
		c.screenX -= 3
	}
	c.boardX--
}

func (c *cursor) moveRight() {
	switch {
	case c.screenX == 28:
		return
	case c.screenX == 8 || c.screenX == 18:
		c.screenX += 4
	default:
		c.screenX += 3
	}
	c.boardX++
}

func (c *cursor) moveUp() {
	switch {
	case c.screenY == 1:
		return
	case c.screenY == 5 || c.screenY == 9:
		c.screenY -= 2
	default:
		c.screenY--
	}
	c.boardY--
}

func (c *cursor) moveDown() {
	switch {
	case c.screenY == 11:
		return
	case c.screenY == 3 || c.screenY == 7:
		c.screenY += 2
	default:
		c.screenY++
	}
	c.boardY++
}
