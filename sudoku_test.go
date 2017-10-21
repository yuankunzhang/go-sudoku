package main

import "testing"

func TestNewFromGridArray(t *testing.T) {
	gridArray := []uint8{
		5, 9, 1, 8, 3, 7, 2, 6, 4,
		6, 4, 2, 1, 9, 5, 3, 8, 7,
		3, 7, 8, 4, 6, 2, 1, 9, 5,
		1, 8, 7, 3, 2, 9, 4, 5, 6,
		9, 3, 5, 7, 4, 6, 8, 1, 2,
		4, 2, 6, 5, 8, 1, 7, 3, 9,
		7, 6, 4, 9, 1, 3, 5, 2, 8,
		8, 1, 9, 2, 5, 4, 6, 7, 3,
		2, 5, 3, 6, 7, 8, 9, 4, 1,
	}

	s, err := NewFromGridArray(gridArray)
	if err != nil {
		t.Error("error creating new sudoku from given array")
	} else {
		if s.Grid[0] != 5 || s.Grid[SudokuNums-1] != 1 {
			t.Error("created a bad sudoku")
		}
	}
}

func TestNewFromGridString(t *testing.T) {
	gridString := ".4..5.3...6.9...52.5..86..7.95..4...2.67.31.5...8..79.5..43..7.31...9.2...8.1..3."

	s, err := NewFromGridString(gridString)
	if err != nil {
		t.Error("error creating new sudoku from given string")
	} else {
		if s.IndexOfNextEmptyCell != 0 {
			t.Error("error to set IndexOfNextEmptyCell")
		}

		if s.Grid[0] != 0 || s.Grid[SudokuNums-2] != 3 {
			t.Error("created a bad sudoku")
		}
	}
}

func TestGetIndexOfRowCell(t *testing.T) {
	var row uint8 = 3
	var cell uint8 = 5
	if getIndexOfRowCell(row, cell) != 32 {
		t.Error("row(3, 5) != index(32)")
	}
}

func TestGetIndexOfColCell(t *testing.T) {
	var col uint8 = 3
	var cell uint8 = 5
	if getIndexOfColCell(col, cell) != 48 {
		t.Error("col(3, 5) != index(48)")
	}
}

func TestGetIndexOfBoxCell(t *testing.T) {
	var box uint8 = 3
	var cell uint8 = 5
	if getIndexOfBoxCell(box, cell) != 38 {
		t.Error("box(3, 5) != index(38)")
	}
}

func TestGetBoxOfIndex(t *testing.T) {
	var idx uint8 = 31
	if GetBoxOfIndex(idx) != 4 {
		t.Error("index(31) != box(4)")
	}
}

func TestIsSolved(t *testing.T) {
	grid1 := []uint8{
		5, 9, 1, 8, 3, 7, 2, 6, 4,
		6, 4, 2, 1, 9, 5, 3, 8, 7,
		3, 7, 8, 4, 6, 2, 1, 9, 5,
		1, 8, 7, 3, 2, 9, 4, 5, 6,
		9, 3, 5, 7, 4, 6, 8, 1, 2,
		4, 2, 6, 5, 8, 1, 7, 3, 9,
		7, 6, 4, 9, 1, 3, 5, 2, 8,
		8, 1, 9, 2, 5, 4, 6, 7, 3,
		2, 5, 3, 6, 7, 8, 9, 4, 1,
	}

	s, _ := NewFromGridArray(grid1)
	if !s.IsSolved() {
		t.Error("grid1 should be solved")
	}

	grid2 := []uint8{
		7, 1, 8, 4, 2, 6, 5, 3, 9,
		2, 6, 5, 3, 8, 9, 4, 7, 1,
		4, 3, 9, 1, 5, 7, 2, 6, 8,
		9, 5, 4, 2, 3, 1, 7, 8, 6,
		1, 8, 2, 7, 6, 4, 3, 9, 5,
		3, 7, 6, 8, 9, 5, 1, 2, 4,
		5, 9, 7, 6, 4, 2, 8, 1, 3,
		6, 2, 3, 5, 1, 8, 9, 4, 7,
		8, 4, 1, 9, 7, 3, 6, 5, 2,
	}

	s, _ = NewFromGridArray(grid2)
	if !s.IsSolved() {
		t.Error("grid2 should be solved")
	}

}
