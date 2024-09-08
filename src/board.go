package src

import (
	"math/rand"
	"time"
)

type cell struct {
	value       uint8
	isGenerated bool
}

type board struct {
	cells                  [9][9]cell
	rows, columns, squares [9]map[uint8]bool
}

func initBoard() *board {
	b := &board{
		cells:   [9][9]cell{},
		rows:    [9]map[uint8]bool{},
		columns: [9]map[uint8]bool{},
		squares: [9]map[uint8]bool{},
	}

	for i := 0; i < 9; i++ {
		b.rows[i] = make(map[uint8]bool)
		b.columns[i] = make(map[uint8]bool)
		b.squares[i] = make(map[uint8]bool)
	}

	return b
}

func generateBoard(difficulty Difficulty) *board {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := initBoard()
	b.fillBoard(0, 0, random)
	b.clearCells(difficulty, random)
	return b
}

func (b *board) fillBoard(row, column int, random *rand.Rand) bool {
	if row == 9 {
		return true
	}

	var available []uint8
	for value := uint8(1); value <= 9; value++ {
		if b.isValid(row, column, value) {
			available = append(available, value)
		}
	}

	ok := false
	for !ok {
		if len(available) == 0 {
			return false
		}

		i := random.Intn(len(available))
		value := available[i]

		b.setCell(row, column, value, true)

		var newRow, newCol int
		if column == 8 {
			newRow, newCol = row+1, 0
		} else {
			newRow, newCol = row, column+1
		}

		ok = b.fillBoard(newRow, newCol, random)
		if !ok {
			available = append(available[:i], available[i+1:]...)
			b.clearCell(row, column, true)
		}
	}
	return true
}

func (b *board) clearCells(difficulty Difficulty, random *rand.Rand) {
	randomCells := random.Perm(81)

	var toRemove int
	switch difficulty {
	case Easy:
		toRemove = 36 + random.Intn(6)
	case Medium:
		toRemove = 42 + random.Intn(8)
	case Hard:
		toRemove = 50 + random.Intn(6)
	case Expert:
		toRemove = 56 + random.Intn(9)
	}

	for toRemove > 0 {
		row := randomCells[toRemove] / 9
		column := randomCells[toRemove] % 9

		b.clearCell(row, column, true)

		toRemove--
	}
}

func (b *board) isValid(row, column int, value uint8) bool {
	square := getSquareFromRowAndColumn(row, column)
	return !b.rows[row][value] && !b.columns[column][value] && !b.squares[square][value]
}

func (b *board) isSolved() bool {
	for i := 0; i < 9; i++ {
		if len(b.rows[i]) != 9 || len(b.columns[i]) != 9 || len(b.squares[i]) != 9 {
			return false
		}
	}
	return true
}

func (b *board) getValue(row, column int) uint8 {
	return b.cells[row][column].value
}

func (b *board) isGenerated(row, column int) bool {
	return b.cells[row][column].isGenerated
}

func (b *board) setCell(row, column int, value uint8, isGenerated bool) {
	if isGenerated || !b.isGenerated(row, column) {
		b.cells[row][column] = cell{value, isGenerated}

		square := getSquareFromRowAndColumn(row, column)
		b.rows[row][value] = true
		b.columns[column][value] = true
		b.squares[square][value] = true
	}
}

func (b *board) clearCell(row, column int, isGenerated bool) {
	if isGenerated || !b.isGenerated(row, column) {
		value := b.getValue(row, column)

		b.cells[row][column] = cell{}

		square := getSquareFromRowAndColumn(row, column)
		delete(b.rows[row], value)
		delete(b.columns[column], value)
		delete(b.squares[square], value)
	}
}

func getSquareFromRowAndColumn(row, column int) int {
	return (row/3)*3 + column/3
}
