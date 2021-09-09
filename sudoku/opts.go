package sudoku

// puzzleOption represents a function that, when called, configures a puzzle.
type puzzleOption func(*Puzzle)

// WithBoxDimensions sets the dimensions of a box in a Sudoku puzzle.
func WithBoxDimensions(height, width PuzzleInt) puzzleOption {
	return func(p *Puzzle) {
		// The product of the height and width should equal the side length of a puzzle.

		// Ensure the dimensions do not exceed the actual puzzle.
		if height > PuzzleInt(len(p.Arr)) {
			height = PuzzleInt(len(p.Arr))
		}
		if width > PuzzleInt(len(p.Arr[0])) {
			width = PuzzleInt(len(p.Arr[0]))
		}
		p.boxHeight, p.boxWidth = height, width
	}
}
