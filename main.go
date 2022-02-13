package main

import (
	"fmt"
	"os"
	"sudoku/board"
	"sudoku/solver"
	"sudoku/validator"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("sudoku takes a csv file and either solves it or validates it")
		fmt.Println("Usage: sudoku-go puzzle.csv")
	}

	args := os.Args[1:]

	fmt.Println(args)

	b := board.New(args[0])

	if b.IsSolved() {
		if validator.Validator(b) {
			fmt.Println("Board is a valid solution")
		} else {
			fmt.Println("There is an error in the board")
		}
	} else {
		if solver.Solver(b) {
			fmt.Println("Solved board!")
			b.PrintBoard()
		} else {
			fmt.Println("The board is not solvable!")
			b.PrintBoard()
		}
	}
}
