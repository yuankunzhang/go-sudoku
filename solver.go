package main

func (s *Sudoku) Solve() bool {
	if s.IsSolved() {
		s.Print()
		return true
	}

	var value uint8
	for value = 1; value <= 9; value++ {
		idx := s.GetIndexOfNextEmptyCell()
		if !s.IsValueValidForIndex(value, idx) {
			continue
		}
		s.Grid[idx] = value
		if s.Solve() {
			return true
		}
		s.Grid[idx] = 0
	}

	return false
}
