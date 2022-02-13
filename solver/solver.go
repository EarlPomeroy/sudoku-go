package solver

import "sudoku/board"

func Solver(b *board.Board) bool {
	for i, row := range b.Board {
		for j, value := range row {
			if value == 0 {
				for k := 1; k <= board.GridSize; k++ {
					if b.IsValid(i, j, uint(k)) {
						b.Board[i][j] = uint(k)
						if Solver(b) {
							return true
						} else {
							b.Board[i][j] = 0
						}
					}
				}

				return false
			}
		}
	}

	return true
}
