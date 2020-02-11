package main

import "fmt"

// Play structure used to play every round
type Play struct {
	value, row, col int
}

// Sudoku game structure that take the game as an argument
type Sudoku struct {
	stack []Play
	board [9][9]int
}

// Finds th next empty cell on the Sudoku boards
func (s *Sudoku) findEmptyCell() (int, int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if s.board[row][col] == 0 {
				return row, col // row, col
			}
		}
	}
	return -1, -1
}

// Try a valid value
func (s *Sudoku) tryValue(row, col int) bool {
	for val := s.board[row][col] + 1; val <= 9; val++ {
		status := false

		if s.checkColumn(col, val) {
			if s.checkRow(row, val) {
				if s.checkBlock(row, col, val) {
					status = true
				}
			}
		}

		if status {
			s.board[row][col] = val
			s.stack = append(s.stack, Play{val, row, col})
			return true
		}
	}
	return false
}

// Check if val
func (s *Sudoku) checkColumn(col, val int) bool {
	for row := 0; row < 9; row++ {
		if s.board[row][col] == val {
			return false
		}
	}
	return true
}

func (s *Sudoku) checkRow(row, val int) bool {
	for col := 0; col < 9; col++ {
		if s.board[row][col] == val {
			return false
		}
	}
	return true
}

func (s *Sudoku) checkCells(rowL, rowH, colL, colH, val int) bool {
	for i := rowL; i < rowH; i++ {
		for j := colL; j < colH; j++ {
			if s.board[i][j] == val {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) checkBlock(row, col, val int) bool {
	if row < 3 {
		if col < 3 {
			return s.checkCells(0, 3, 0, 3, val)
		} else if 3 <= col && col <= 5 {
			return s.checkCells(0, 3, 3, 6, val)
		} else if 5 < col {
			return s.checkCells(0, 3, 6, 9, val)
		}
	} else if 3 <= row && row <= 5 {
		if col < 3 {
			return s.checkCells(3, 6, 0, 3, val)
		} else if 3 <= col && col <= 5 {
			return s.checkCells(3, 6, 3, 6, val)
		} else if 5 < col {
			return s.checkCells(3, 6, 6, 9, val)
		}
	} else if 5 < row {
		if col < 3 {
			return s.checkCells(6, 9, 0, 3, val)
		} else if 3 <= col && col <= 5 {
			return s.checkCells(6, 9, 3, 6, val)
		} else if 5 < col {
			return s.checkCells(6, 9, 6, 9, val)
		}
	}
	return false
}

func (s *Sudoku) backtrack() {
	backtrackStatus := false
	for !backtrackStatus {
		var val Play
		val, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
		backtrackStatus = s.tryValue(val.row, val.col)
		if !backtrackStatus {
			s.board[val.row][val.col] = 0
		}
	}
}

func (s *Sudoku) TestSolve() {
	row, col := 0, 0

	for row != -1 {
		// find an emty cell
		row, col = s.findEmptyCell()
		// try a value in this cell
		if row != -1 {
			if s.tryValue(row, col) {
				continue
			} else {
				s.backtrack()
			}
		} else {
			break
		}
	}

}

func (s *Sudoku) printBoard() {
	for i := 0; i < 9; i++ {
		if (i%3) == 0 && i != 0 {
			fmt.Println("- - - - - - - - - - -")
		}

		for j := 0; j < 9; j++ {
			if (j%3) == 0 && j != 0 {
				fmt.Printf("| ")
			}
			if j == 8 {
				fmt.Printf("%d\n", s.board[i][j])
			} else {
				fmt.Printf("%d ", s.board[i][j])
			}
		}
	}
	fmt.Printf("\n")
}

func main() {
	theGame := [9][9]int{{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0}}
	game := Sudoku{board: theGame}
	game.printBoard()
	game.TestSolve()
	game.printBoard()
}
