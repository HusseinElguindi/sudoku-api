package sudoku

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		puzzle [][]PuzzleInt
		solved [][]PuzzleInt
	}{
		{
			puzzle: [][]PuzzleInt{
				{8, 0, 0, 4, 0, 0, 9, 1, 0},
				{0, 0, 3, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 3, 0, 0, 4},
				{0, 0, 0, 0, 0, 1, 0, 4, 0},
				{0, 5, 8, 0, 0, 0, 7, 0, 0},
				{0, 7, 0, 0, 0, 6, 8, 0, 0},
				{0, 0, 0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 6, 0},
				{9, 1, 0, 0, 6, 0, 5, 0, 0},
			},
			solved: [][]PuzzleInt{
				{8, 6, 5, 4, 2, 7, 9, 1, 3},
				{2, 4, 3, 9, 1, 5, 6, 8, 7},
				{7, 9, 1, 6, 8, 3, 2, 5, 4},
				{6, 2, 9, 8, 7, 1, 3, 4, 5},
				{1, 5, 8, 3, 4, 9, 7, 2, 6},
				{3, 7, 4, 2, 5, 6, 8, 9, 1},
				{5, 8, 6, 1, 3, 2, 4, 7, 9},
				{4, 3, 7, 5, 9, 8, 1, 6, 2},
				{9, 1, 2, 7, 6, 4, 5, 3, 8},
			},
		},
		{
			puzzle: [][]PuzzleInt{
				{8, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 3, 6, 0, 0, 0, 0, 0},
				{0, 7, 0, 0, 9, 0, 2, 0, 0},
				{0, 5, 0, 0, 0, 7, 0, 0, 0},
				{0, 0, 0, 0, 4, 5, 7, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 3, 0},
				{0, 0, 1, 0, 0, 0, 0, 6, 8},
				{0, 0, 8, 5, 0, 0, 0, 1, 0},
				{0, 9, 0, 0, 0, 0, 4, 0, 0},
			},
			solved: [][]PuzzleInt{
				{8, 1, 2, 7, 5, 3, 6, 4, 9},
				{9, 4, 3, 6, 8, 2, 1, 7, 5},
				{6, 7, 5, 4, 9, 1, 2, 8, 3},
				{1, 5, 4, 2, 3, 7, 8, 9, 6},
				{3, 6, 9, 8, 4, 5, 7, 2, 1},
				{2, 8, 7, 1, 6, 9, 5, 3, 4},
				{5, 2, 1, 9, 7, 4, 3, 6, 8},
				{4, 3, 8, 5, 2, 6, 9, 1, 7},
				{7, 9, 6, 3, 1, 8, 4, 5, 2},
			},
		},
	}
	for _, tc := range testCases {
		puzzle := NewPuzzle(tc.puzzle)
		require.True(t, puzzle.Solve())
		require.Equal(t, tc.solved, tc.puzzle)
	}
}

func TestString(t *testing.T) {
	arr := make([][]PuzzleInt, 9)
	for i := range arr {
		arr[i] = make([]PuzzleInt, 9)
	}
	puzzle := NewPuzzle(arr)
	require.Equal(t, puzzle.String(), "[[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0]]")
}

func TestRowContains(t *testing.T) {
	type test struct {
		row, val PuzzleInt
		expected bool
	}
	testCases := []struct {
		arr   [][]PuzzleInt
		tests []test
	}{
		{
			arr: [][]PuzzleInt{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 5, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{9, 0, 1, 0, 0, 0, 0, 0, 0},
			},
			tests: []test{
				{8, 9, true},
				{8, 1, true},
				{8, 2, false},
				{4, 5, true},
			},
		},
	}
	for _, tc := range testCases {
		puzzle := NewPuzzle(tc.arr)
		for _, test := range tc.tests {
			require.Equal(t, test.expected, puzzle.rowContains(test.row, test.val))
		}
	}
}

func TestColContains(t *testing.T) {
	type test struct {
		col, val PuzzleInt
		expected bool
	}
	testCases := []struct {
		arr   [][]PuzzleInt
		tests []test
	}{
		{
			arr: [][]PuzzleInt{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 5, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{9, 0, 1, 0, 0, 0, 0, 0, 0},
			},
			tests: []test{
				{4, 5, true},
				{4, 1, false},
				{0, 9, true},
				{1, 1, false},
			},
		},
	}
	for _, tc := range testCases {
		puzzle := NewPuzzle(tc.arr)
		for _, test := range tc.tests {
			require.Equal(t, test.expected, puzzle.colContains(test.col, test.val))
		}
	}
}

func TestBoxContains(t *testing.T) {
	type test struct {
		row, col, val PuzzleInt
		expected      bool
	}
	testCases := []struct {
		arr   [][]PuzzleInt
		tests []test
	}{
		{
			arr: [][]PuzzleInt{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 5, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{9, 0, 1, 0, 0, 0, 0, 0, 0},
			},
			tests: []test{
				{6, 2, 9, true},
				{6, 2, 1, true},
				{6, 2, 5, false},
				{4, 4, 5, true},
				{4, 4, 4, false},
			},
		},
	}
	for _, tc := range testCases {
		puzzle := NewPuzzle(tc.arr)
		for _, test := range tc.tests {
			require.Equal(t, test.expected, puzzle.boxContains(test.row, test.col, test.val))
		}
	}
}

func TestValidPos(t *testing.T) {
	type test struct {
		row, col, val PuzzleInt
		expected      bool
	}
	testCases := []struct {
		arr   [][]PuzzleInt
		tests []test
	}{
		{
			arr: [][]PuzzleInt{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 5, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0},
			},
			tests: []test{
				{9, 9, 1, false}, // Out of bounds
				{4, 4, 5, false}, // Position is not vacant
				{4, 3, 5, false}, // Duplicate digit on row
				{3, 4, 5, false}, // Duplicate digit on col
				{5, 5, 5, false}, // Duplicate digit in box
				{0, 0, 1, true},
			},
		},
	}
	for _, tc := range testCases {
		puzzle := NewPuzzle(tc.arr)
		for _, test := range tc.tests {
			require.Equal(t, test.expected, puzzle.isValidPos(test.row, test.col, test.val))
		}
	}
}

func TestNextEmptyPos(t *testing.T) {
	type test struct {
		expectRow, expectCol PuzzleInt
		expectedOk           bool
	}
	testCases := []struct {
		arr   [][]PuzzleInt
		tests []test
	}{
		{
			arr: [][]PuzzleInt{
				{2, 2, 2},
				{2, 0, 2},
				{2, 2, 0},
			},
			tests: []test{
				{1, 1, true},
				{2, 2, true},
				{0, 0, false},
			},
		},
	}
	for _, tc := range testCases {
		puzzle := NewPuzzle(tc.arr)
		var startRow, startCol PuzzleInt
		for _, test := range tc.tests {
			row, col, ok := puzzle.nextEmptyPos(startRow, startCol)
			require.Equal(t, test.expectedOk, ok)
			if ok {
				require.Equal(t, test.expectRow, row)
				require.Equal(t, test.expectCol, col)
				puzzle.Arr[row][col] = 2 // Occupy position
			}
			startRow, startCol = test.expectRow, test.expectCol
		}
	}
}
