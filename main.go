package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"sudoku/src"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: sudoku [ easy | medium | hard ]")
		os.Exit(1)
	}

	var difficulty src.Difficulty
	switch os.Args[1] {
	case "easy":
		difficulty = src.Easy
	case "medium":
		difficulty = src.Medium
	case "hard":
		difficulty = src.Hard
	default:
		fmt.Println("Usage: sudoku [ easy | medium | hard ]")
		os.Exit(1)
	}

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	src.NewGame(difficulty).Run()
}
