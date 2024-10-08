package main

import (
	"math/rand"
	"time"
)

type cell struct {
	value       uint8
	isGenerated bool
	isValid     bool
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

	var toRemove int
	switch difficulty {
	case Easy:
		toRemove = 36 + random.Intn(6)
	case Medium:
		toRemove = 42 + random.Intn(8)
	case Hard:
		toRemove = 50 + random.Intn(6)
	}
	b.clearCells(toRemove, 0, random.Perm(81))

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

		newRow, newColumn := getNewRowAndColumn(row, column)

		ok = b.fillBoard(newRow, newColumn, random)
		if !ok {
			available = append(available[:i], available[i+1:]...)
			b.clearCell(row, column, true)
		}
	}
	return true
}

func (b *board) clearCells(toRemove, i int, positions []int) {
	if toRemove == 0 {
		return
	}

	row := positions[i] / 9
	column := positions[i] % 9
	value := b.getCell(row, column).value

	b.clearCell(row, column, true)

	if b.countSolutions(0, 0) != 1 {
		b.setCell(row, column, value, true)
		b.clearCells(toRemove, i+1, positions)
	} else {
		b.clearCells(toRemove-1, i+1, positions)
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

func (b *board) getCell(row, column int) cell {
	return b.cells[row][column]
}

func (b *board) setCell(row, column int, value uint8, isGenerated bool) (isValid bool) {
	if isGenerated {
		if b.isValid(row, column, value) {
			isValid = true
			b.cells[row][column] = cell{value, true, true}

			square := getSquareFromRowAndColumn(row, column)
			b.rows[row][value] = true
			b.columns[column][value] = true
			b.squares[square][value] = true
		} else {
			isValid = false
		}
	} else if !b.getCell(row, column).isGenerated {
		isValid = b.isValid(row, column, value)
		b.cells[row][column] = cell{value, false, isValid}

		if isValid {
			square := getSquareFromRowAndColumn(row, column)
			b.rows[row][value] = true
			b.columns[column][value] = true
			b.squares[square][value] = true
		} else {
			mistakes++
		}
	}
	return
}

func (b *board) clearCell(row, column int, isGenerated bool) {
	if isGenerated || !b.getCell(row, column).isGenerated {
		value := b.getCell(row, column).value
		isValid := b.getCell(row, column).isValid

		b.cells[row][column] = cell{}

		if isValid {
			square := getSquareFromRowAndColumn(row, column)
			delete(b.rows[row], value)
			delete(b.columns[column], value)
			delete(b.squares[square], value)
		}
	}
}

func (b *board) countSolutions(row, column int) (count int) {
	if row == 9 {
		return 1
	} else if b.getCell(row, column).value != 0 {
		return b.countSolutions(getNewRowAndColumn(row, column))
	}

	for value := uint8(1); value <= 9; value++ {
		if b.isValid(row, column, value) {
			b.setCell(row, column, value, false)

			count += b.countSolutions(getNewRowAndColumn(row, column))

			b.clearCell(row, column, false)
		}
	}
	return
}

func getSquareFromRowAndColumn(row, column int) int {
	return (row/3)*3 + column/3
}

func getNewRowAndColumn(row, column int) (newRow, newColumn int) {
	newRow, newColumn = row, column+1
	if newColumn == 9 {
		newRow, newColumn = row+1, 0
	}
	return
}
