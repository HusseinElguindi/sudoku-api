/*
Package sudoku implements a flexible recursive backtracking Sudoku puzzle solver library.

The recommended usage is:
	arr := [][]pInt{...}
	puzzle := sudoku.NewPuzzle(arr) // arr should not be edited anymore
	fmt.Println(puzzle.Pretty())
	puzzle.Solve()
	fmt.Println(puzzle.Pretty())
*/
package sudoku

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
)

// TODO: rethink the whole hVal, vVal system
// it is not needed, we can keep a set for the horizontal that we're working on

var (
	ErrInvalidPuzzle = errors.New("invalid/unsolvable Sudoku puzzle")
)

type PuzzleInt = uint16
type Puzzle struct {
	Arr                 [][]PuzzleInt
	boxHeight, boxWidth PuzzleInt

	// Acts as a hashset for the values in a row (hVals) or column (vVals).
	// Could use bit manipulation, []int.
	// hVals, vVals [][]bool
}

func NewPuzzle(arr [][]PuzzleInt, opts ...puzzleOption) Puzzle {
	// TODO: compare puzzle array size with max of PuzzleInt

	// Validate number of rows
	if len(arr) == 0 {
		panic("invalid number of puzzle rows (0)")
	}
	// Validate number of columns
	cols := len(arr[0])
	if cols == 0 {
		panic("invalid number of puzzle columns (0)")
	}
	// Ensure all rows have the same number of columns
	for _, row := range arr {
		if len(row) != cols {
			panic("all rows must have a uniform number of columns")
		}
	}

	puzzle := Puzzle{
		Arr: arr,
		// Box dimensions are defaulted to the square root of each puzzle side.
		boxHeight: PuzzleInt(math.Sqrt(float64(len(arr)))),
		boxWidth:  PuzzleInt(math.Sqrt(float64(cols))),
	}

	// Apply options
	for _, opt := range opts {
		opt(&puzzle)
	}

	// Ensure non-zero box dimensions for division
	if puzzle.boxHeight == 0 {
		puzzle.boxHeight = PuzzleInt(len(puzzle.Arr))
	}
	if puzzle.boxWidth == 0 {
		puzzle.boxWidth = PuzzleInt(cols)
	}

	return puzzle
}

// rowContains returns whether or not the row contains val.
func (p Puzzle) rowContains(row, val PuzzleInt) bool {
	for _, n := range p.Arr[row] {
		if val == n {
			return true
		}
	}
	return false
}

// colContains returns whether or not the col contains val.
func (p Puzzle) colContains(col, val PuzzleInt) bool {
	for row := range p.Arr {
		if p.Arr[row][col] == val {
			return true
		}
	}
	return false
}

// boxContains returns whether or not the box contains val.
func (p Puzzle) boxContains(row, col, val PuzzleInt) bool {
	// Calculate the current box's coordinates (top left position).
	boxRow, boxCol := (row/p.boxWidth)*p.boxWidth, (col/p.boxHeight)*p.boxHeight

	// Iterate over all cells in the current box.
	for row = 0; row < p.boxHeight; row++ {
		for col = 0; col < p.boxWidth; col++ {
			if p.Arr[boxRow+row][boxCol+col] == val {
				return true
			}
		}
	}
	return false
}

// isValidPos returns whether or not the value at the row and col position follows the
// constraints of the Sudoku puzzle. The constraints are that a digit can only appear once
// in a row, column, and box. In addition, a position must be vacant.
func (p Puzzle) isValidPos(row, col, val PuzzleInt) bool {
	return p.Arr[row][col] == 0 && // The position is already occupied
		!p.rowContains(row, val) && // Another position on the row has the same value
		!p.colContains(col, val) && // Another position on the column has the same value
		!p.boxContains(row, col, val) // Another position in the box has the same value
}

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
	// Calculate the number of horizontal lines
	hLines := int(math.Max(0, float64(len(p.Arr)/int(p.boxWidth)-1)))

	for i, row := range p.Arr {
		// Place horizontal lines for every box height
		if i%int(p.boxHeight) == 0 && i != 0 {
			for i := 0; i < len(p.Arr)+hLines; i++ {
				sb.WriteString("- ")
			}
			sb.WriteByte('\n')
		}
		for j, v := range row {
			// Place a horizontal line every box width
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

// nextEmptyPos returns the next unoccupied position of the puzzle. If there are none, ok is false.
// The search begins from the passed start row and col position, which is inclusive.
func (p Puzzle) nextEmptyPos(startRow, startCol PuzzleInt) (row, col PuzzleInt, ok bool) {
	row, col = startRow, startCol
	for ; row < PuzzleInt(len(p.Arr)); row++ {
		for ; col < PuzzleInt(len(p.Arr[0])); col++ {
			if p.Arr[row][col] == 0 {
				return row, col, true
			}
		}
		col = 0 // Reset col from startCol
	}
	return
}

func (p Puzzle) Solve() bool {
	// TODO: iterate over every point to validate puzzle before solving.
	return p.solve(0, 0)
}

func (p Puzzle) solve(row, col PuzzleInt) bool {
	// Find the next empty position, if any.
	row, col, ok := p.nextEmptyPos(row, col)
	if !ok {
		return true
	}

	// Try all possible values, recurse, and backtrack.
	// Can be optimized using memoization via sets.
	// TODO: Puzzle must always be a square?
	for val := PuzzleInt(1); val <= PuzzleInt(len(p.Arr)); val++ {
		if !p.isValidPos(row, col, val) {
			continue
		}
		p.Arr[row][col] = val
		if ok := p.solve(row, col); ok {
			return ok
		}
		// Reset position value next attempt (backtrack)
		p.Arr[row][col] = 0
	}

	// Already attempted all possible values for this position.
	return false
}
