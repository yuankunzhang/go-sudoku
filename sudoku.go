package main

import (
	"errors"
	"fmt"
)

const SudokuSize = 9
const SudokuNums = SudokuSize * SudokuSize

// Sudoku represents a sudoku state.
type Sudoku struct {
	IndexOfNextEmptyCell int
	Grid                 [SudokuNums]uint8
}

// NewFromGridArray creates a new sudoku from given array.
func NewFromGridArray(gridArray []uint8) (*Sudoku, error) {
	if len(gridArray) != SudokuNums {
		return nil, errors.New("invalid grid array length")
	}

	s := &Sudoku{IndexOfNextEmptyCell: SudokuNums}

	for i := 0; i < SudokuNums; i++ {
		if gridArray[i] < 0 || gridArray[i] > 9 {
			return nil, errors.New("grid array contains invalid value(s)")
		}
		s.Grid[i] = gridArray[i]
		if gridArray[i] == 0 && s.IndexOfNextEmptyCell == SudokuNums {
			s.IndexOfNextEmptyCell = i
		}
	}

	return s, nil
}

// NewFromGridString creates a new sudoku from given string.
func NewFromGridString(gridString string) (*Sudoku, error) {
	if len(gridString) != SudokuNums {
		return nil, errors.New("invalid grid string length")
	}

	gridArray := make([]uint8, SudokuNums)

	for i := 0; i < SudokuNums; i++ {
		c := gridString[i]
		if c == '.' || c == '0' {
			gridArray[i] = 0
		} else if c > '0' || c < '9' {
			gridArray[i] = c - '0'
		} else {
			return nil, errors.New("grid string contains invalid value(s)")
		}
	}

	return NewFromGridArray(gridArray)
}

func getIndexOfRowCell(row, cell uint8) uint8 {
	return row*9 + cell
}

func getIndexOfColCell(col, cell uint8) uint8 {
	return cell*9 + col
}

func getIndexOfBoxCell(box, cell uint8) uint8 {
	row, col := box/3*3+cell/3, box%3*3+cell%3
	return row*9 + col
}

func getRowOfIndex(idx uint8) uint8 {
	return idx / 9
}

func getColOfIndex(idx uint8) uint8 {
	return idx % 9
}

func GetBoxOfIndex(idx uint8) uint8 {
	row, col := idx/9, idx%9
	return row/3*3 + col/3
}

func (s *Sudoku) isValueValidForRow(value, row uint8) bool {
	for cell := 0; cell < SudokuSize; cell++ {
		idx := getIndexOfRowCell(row, uint8(cell))
		if value == s.Grid[idx] {
			return false
		}
	}
	return true
}

func (s *Sudoku) isValueValidForCol(value, col uint8) bool {
	for cell := 0; cell < SudokuSize; cell++ {
		idx := getIndexOfColCell(col, uint8(cell))
		if value == s.Grid[idx] {
			return false
		}
	}
	return true
}

func (s *Sudoku) isValueValidForBox(value, box uint8) bool {
	for cell := 0; cell < SudokuSize; cell++ {
		idx := getIndexOfBoxCell(box, uint8(cell))
		if value == s.Grid[idx] {
			return false
		}
	}
	return true
}

func (s *Sudoku) IsValueValidForIndex(value, idx uint8) bool {
	row := getRowOfIndex(idx)
	if !s.isValueValidForRow(value, row) {
		return false
	}

	col := getColOfIndex(idx)
	if !s.isValueValidForCol(value, col) {
		return false
	}

	box := GetBoxOfIndex(idx)
	if !s.isValueValidForBox(value, box) {
		return false
	}

	return true
}

func (s *Sudoku) IsSolved() bool {
	for i := 0; i < SudokuSize; i++ {
		rowValues := []uint8{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		colValues := []uint8{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		boxValues := []uint8{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		for j := 0; j < SudokuSize; j++ {
			// Row.
			idx := getIndexOfRowCell(uint8(i), uint8(j))
			if rowValues[s.Grid[idx]] == 1 {
				return false
			}
			rowValues[s.Grid[idx]] = 1

			// Col.
			idx = getIndexOfColCell(uint8(i), uint8(j))
			if colValues[s.Grid[idx]] == 1 {
				return false
			}
			colValues[s.Grid[idx]] = 1

			// Box.
			idx = getIndexOfBoxCell(uint8(i), uint8(j))
			if boxValues[s.Grid[idx]] == 1 {
				return false
			}
			boxValues[s.Grid[idx]] = 1
		}
	}
	return true
}

func (s *Sudoku) GetIndexOfNextEmptyCell() uint8 {
	var idx uint8
	for idx = 0; idx < SudokuNums; idx++ {
		if s.Grid[idx] == 0 {
			break
		}
	}
	return idx
}

func (s *Sudoku) Print() {
	for i := 0; i < SudokuSize; i++ {
		for j := 0; j < SudokuSize; j++ {
			c := '0' + s.Grid[i*9+j]
			if c == '0' {
				c = '.'
			}
			fmt.Printf("%c ", c)
		}
		fmt.Println()
	}
	fmt.Println()
}
