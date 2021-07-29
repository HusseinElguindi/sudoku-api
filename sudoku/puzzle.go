/*
Package sudoku implements a flexible recursive backtracking Sudoku puzzle solver library.

The recommended usage is:
	arr := [][]PuzzleInt{...}
	puzzle := sudoku.NewPuzzle(arr) // arr should not be edited anymore

	if solved := puzzle.Solve(); solved {
		fmt.Println("Solved!")
	} else {
		fmt.Println("Could not solve (invalid puzzle).")
	}
	// Read the value of puzzle.Arr or arr after solve.
	...
	fmt.Println(puzzle.Pretty())
*/
package sudoku

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

// PuzzleInt reprents the integer type used in the puzzle.
type PuzzleInt = uint16

// Puzzle represents a Sudoku puzzle with an underlying matrix.
type Puzzle struct {
	Arr                 [][]PuzzleInt
	boxHeight, boxWidth PuzzleInt

	// Acts as a bitset for the values in a row (rowVals), column (colVals), or box (boxVals).
	rowVals, colVals []bitSet
	boxVals          [][]bitSet
}

// NewPuzzle constructs a new puzzle with the passed matrix. Options may be passed
// to further configure a puzzle. arr should not be altered after puzzle creation as
// it is used as the underlying array for the puzzle. New puzzles should not reuse
// an old Puzzle struct, but should construct a new Puzzle.
func NewPuzzle(arr [][]PuzzleInt, opts ...puzzleOption) Puzzle {
	// TODO: compare puzzle array size with max of PuzzleInt.

	// Validate number of rows.
	rows := len(arr)
	if rows == 0 {
		panic("invalid number of puzzle rows (0)")
	}
	// Validate number of columns.
	cols := len(arr[0])
	if cols == 0 {
		panic("invalid number of puzzle columns (0)")
	}
	// Ensure the puzzle is a square.
	if rows != cols {
		panic("invalid puzzle, must be a square")
	}
	// Ensure all rows have the same number of columns.
	for _, row := range arr {
		if len(row) != cols {
			panic("all rows must have a uniform number of columns")
		}
	}

	puzzle := Puzzle{
		Arr: arr,
		// Box dimensions default to the square root of each puzzle side.
		boxHeight: PuzzleInt(math.Sqrt(float64(rows))),
		boxWidth:  PuzzleInt(math.Sqrt(float64(cols))),

		rowVals: make([]bitSet, rows),
		colVals: make([]bitSet, cols),
	}

	// Apply options.
	for _, opt := range opts {
		opt(&puzzle)
	}

	// Ensure non-zero box dimensions for division.
	if puzzle.boxHeight == 0 {
		puzzle.boxHeight = PuzzleInt(rows)
	}
	if puzzle.boxWidth == 0 {
		puzzle.boxWidth = PuzzleInt(cols)
	}

	// Calculate number of boxes and its bitset.
	// rows and cols should be divisible by boxHeight and boxWidth, respectively.
	puzzle.boxVals = make([][]bitSet, rows/int(puzzle.boxHeight))
	for boxRow := range puzzle.boxVals {
		puzzle.boxVals[boxRow] = make([]bitSet, cols/int(puzzle.boxWidth))
	}

	return puzzle
}

// TODO: next 3 constraint functions have a pattern (like: get bitset and populate func)
// see the pattern and make a file "constraints.go" where code can be reused

// rowContains returns whether or not the row contains val.
func (p Puzzle) rowContains(row, val PuzzleInt) bool {
	bs := &p.rowVals[row]
	// Populate the bitset if empty or val is 0 (0 is the value replaced during a solve).
	if bs.Len() == 0 || val == 0 {
		for _, n := range p.Arr[row] {
			bs.Set(int(n), 1)
		}
	}
	// Return true if val is found in the set.
	return bs.Get(int(val)) == 1
}

// colContains returns whether or not the col contains val.
func (p Puzzle) colContains(col, val PuzzleInt) bool {
	bs := &p.colVals[col]
	// Populate the bitset if empty.
	if bs.Len() == 0 || val == 0 {
		for _, row := range p.Arr {
			bs.Set(int(row[col]), 1)
		}
	}
	// Return true if val is found in the set.
	return bs.Get(int(val)) == 1
}

// boxIndex calculates the index of the current box.
func (p Puzzle) boxIndex(row, col PuzzleInt) (boxRow, boxCol PuzzleInt) {
	return row / p.boxHeight, col / p.boxWidth
}

// boxContains returns whether or not the box contains val.
func (p Puzzle) boxContains(row, col, val PuzzleInt) bool {
	// Calculate the box's bitset index.
	boxRow, boxCol := p.boxIndex(row, col)
	bs := &p.boxVals[boxRow][boxCol]

	// If the box's bitset is empty, populate it.
	if bs.Len() == 0 || val == 0 {
		// Calculate the box's puzzle coordinates (top left position).
		boxRow, boxCol := boxRow*p.boxHeight, boxCol*p.boxWidth

		// Iterate over all cells in the current box.
		for row = 0; row < p.boxHeight; row++ {
			for col = 0; col < p.boxWidth; col++ {
				bs.Set(int(p.Arr[boxRow+row][boxCol+col]), 1)
			}
		}
	}
	// Return true if val is found in the set.
	return bs.Get(int(val)) == 1
}

// isValidPos returns whether or not the value at the row and col position follows the
// constraints of the Sudoku puzzle. The constraints are that a digit can only appear once
// in a row, column, and box. In addition, a position must be vacant.
func (p Puzzle) isValidPos(row, col, val PuzzleInt) bool {
	return int(row) < len(p.Arr) && int(col) < len(p.Arr[0]) && // Bounds check
		p.Arr[row][col] == 0 && // The position must be vacant
		!p.rowContains(row, val) && // Another row position should not have the same value
		!p.colContains(col, val) && // Another column position should not have the same value
		!p.boxContains(row, col, val) // Another box position should not have the same value
}

// nextEmptyPos returns the next unoccupied position of the puzzle. If there are none, ok is false.
// The search begins from the passed start row and col position, which is inclusive.
func (p Puzzle) nextEmptyPos(startRow, startCol PuzzleInt) (row, col PuzzleInt, ok bool) {
	row, col = startRow, startCol
	for ; row < PuzzleInt(len(p.Arr)); row++ {
		for ; col < PuzzleInt(len(p.Arr[row])); col++ {
			if p.Arr[row][col] == 0 {
				return row, col, true
			}
		}
		col = 0 // Reset col from startCol after first iteration
	}
	return
}

// Solve solves modifies the underlying array to solve the Sudoku puzzle recursively, backtracking
// when an invalid value is guessed, until a solution is found. Solve returns true when a puzzle is
// successfully solved, otherwise, the puzzle was unsolvable.
func (p Puzzle) Solve() bool {
	// TODO: validate puzzle before solving.
	return p.solve(0, 0)
}
func (p Puzzle) solve(row, col PuzzleInt) bool {
	// Find the next empty position, if any.
	row, col, ok := p.nextEmptyPos(row, col)
	if !ok {
		// No empty position, the puzzle is solved if it is valid.
		return true
	}

	// Try all possible values, recurse, and backtrack.
	for val := PuzzleInt(1); val <= PuzzleInt(len(p.Arr)); val++ {
		// Validate position for the current value.
		if !p.isValidPos(row, col, val) {
			continue
		}
		// Set the value.
		p.Arr[row][col] = val
		// Update bitsets.
		boxRow, boxCol := p.boxIndex(row, col)
		p.boxVals[boxRow][boxCol].Set(int(val), 1)
		p.rowVals[row].Set(int(val), 1)
		p.colVals[col].Set(int(val), 1)

		// Try to solve this path by recursing, return if that path was successful.
		if ok := p.solve(row, col); ok {
			return true
		}
		// The path was not successful, reset position (backtrack).
		p.Arr[row][col] = 0
		// Reset bitsets.
		p.boxVals[boxRow][boxCol].Set(int(val), 0)
		p.rowVals[row].Set(int(val), 0)
		p.colVals[col].Set(int(val), 0)
	}

	// Already attempted all possible values for this position.
	return false
}

// String implements the Stringer interface for Puzzle by encoding into JSON.
func (p Puzzle) String() string {
	b, err := json.Marshal(p.Arr)
	if err != nil {
		panic("could not decode puzzle into json")
	}
	return string(b)
}

// Pretty returns a formatted string representation of the Sudoku puzzle for human
// readability.
func (p Puzzle) Pretty() string {
	sb := strings.Builder{}
	// Calculate the number of horizontal lines.
	hLines := int(math.Max(0, float64(len(p.Arr)/int(p.boxWidth)-1)))

	for i, row := range p.Arr {
		// Place horizontal lines for every box height.
		if i%int(p.boxHeight) == 0 && i != 0 {
			for i := 0; i < len(p.Arr)+hLines; i++ {
				sb.WriteString("- ")
			}
			sb.WriteByte('\n')
		}
		for j, v := range row {
			// Place a horizontal line every box width.
			if j%int(p.boxWidth) == 0 && j != 0 {
				sb.WriteString("| ")
			}
			if v == 0 {
				sb.WriteByte(' ')
			} else {
				sb.WriteString(fmt.Sprintf("%d", v))
			}
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
