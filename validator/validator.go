package validator

import "sudoku/board"

func Validator(b *board.Board) bool {
	for i, row := range b.Board {
		for j, value := range row {
			if !b.IsValid(i, j, value) {
				return false
			}
		}
	}

	return true
}
