# Sudoku Validator/Solver
___
## The purpose of this is to validate or solve a CSV file containing a 9x9 grid of number and verify if they are a solution for a Sudoku puzzle.

The application will detect if the puzzle is solved by looking for blanks in the "board". Blanks are represented in the CSV file by the number 0.

No special libraries were used.

###To build debug, run for the root of project

`$ go build`

###To build release, run for the root of project
`$ go build -ldflags "-s -w"`

###To run the project

`$ go run main.go samples/example_valid.csv`

or

`$ go run main.go samples/puzzle1.csv`