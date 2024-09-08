package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: sudoku [ easy | medium | hard ]")
		os.Exit(1)
	}

	var difficulty Difficulty
	switch os.Args[1] {
	case "easy":
		difficulty = Easy
	case "medium":
		difficulty = Medium
	case "hard":
		difficulty = Hard
	default:
		fmt.Println("Usage: sudoku [ easy | medium | hard ]")
		os.Exit(1)
	}

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	NewGame(difficulty).Run()
}
