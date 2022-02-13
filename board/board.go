package board

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const GridSize = 9

type Board struct {
	Board [GridSize][GridSize]uint
}

func New(filename string) *Board {
	b := new(Board)
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for row := 0; row < GridSize; row++ {
		for column := 0; column < GridSize; column++ {

			intVal, err := strconv.ParseUint(data[row][column], 10, 32)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			b.Board[row][column] = uint(intVal)
		}
	}

	return b
}

func (b *Board) IsSolved() bool {
	for _, row := range b.Board {
		for _, col := range row {
			if col == 0 {
				return false
			}
		}
	}

	return true
}

func (b *Board) PrintBoard() {
	for i, row := range b.Board {
		if i > 0 && i%3 == 0 {
			fmt.Println("===========")
		}

		for j, col := range row {
			if j > 0 && j%3 == 0 {
				if j < GridSize {
					fmt.Print("|")
				}
			}
			fmt.Print(col)
		}

		fmt.Println()
	}
}

func (b *Board) IsValid(row, col int, value uint) bool {
	return b.validForRow(row, col, value) && b.validForColumn(row, col, value) && b.validForSquare(row, col, value)
}

func (b *Board) validForRow(row, col int, value uint) bool {
	for i, c := range b.Board[row] {
		if i != col {
			if c == value {
				return false
			}
		}
	}

	return true
}

func (b *Board) validForColumn(row, col int, value uint) bool {
	for i, r := range b.Board {
		if i != row {
			if r[col] == value {
				return false
			}
		}
	}

	return true
}

func (b *Board) validForSquare(row, col int, value uint) bool {
	leftMostRow := row - row%3
	topMostColumn := col - col%3

	for r := leftMostRow; r < leftMostRow+3; r++ {
		for c := topMostColumn; c < topMostColumn+3; c++ {
			if row != r || col != c {
				if b.Board[r][c] == value {
					return false
				}
			}
		}
	}

	return true
}
