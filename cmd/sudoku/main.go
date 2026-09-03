package main

import (
	"fmt"
	"github.com/MhcVintar/sudoku/internal/game"
	"github.com/nsf/termbox-go"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: sudoku [ easy | medium | hard ]")
		os.Exit(1)
	}

	var difficulty game.Difficulty
	switch os.Args[1] {
	case "easy":
		difficulty = game.Easy
	case "medium":
		difficulty = game.Medium
	case "hard":
		difficulty = game.Hard
	default:
		fmt.Println("Usage: sudoku [ easy | medium | hard ]")
		os.Exit(1)
	}

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	game.NewGame(difficulty).Run()
}
