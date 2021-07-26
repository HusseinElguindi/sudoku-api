package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

// createRandomPuzzle generates a random puzzle and inserts it into the testQueries database.
// The inserted puzzle is returned and guaranteed to be valid. Otherwise, the test (t) is failed.
func createRandomPuzzle(t *testing.T) Puzzle {
	// Generate a random string of 25 ASCII letters
	arrayStr := gofakeit.LetterN(25)
	// Insert puzzle into db
	puzzle, err := testQueries.CreatePuzzle(context.Background(), arrayStr)

	// Validate the inserted values
	require.NoError(t, err)
	require.NotEmpty(t, puzzle)
	require.NotZero(t, puzzle.ID)

	// Compare the inserted values with the original values
	require.Equal(t, puzzle.ArrayStr, arrayStr)
	require.WithinDuration(t, puzzle.CreatedAt, time.Now(), time.Second)

	return puzzle
}

// TestCreatePuzzle inserts a random puzzle into the global test db, failing the test on any errors
// or unexpected results.
func TestCreatePuzzle(t *testing.T) { createRandomPuzzle(t) }

// TestGetPuzzleByID inserts a random puzzle into the db and attempts to query it by its ID.
// The test fails if no puzzle was returned or it does not match the inserted puzzle.
func TestGetPuzzleByID(t *testing.T) {
	insertPuzzle := createRandomPuzzle(t)
	getPuzzle, err := testQueries.GetPuzzleByID(context.Background(), insertPuzzle.ID)

	require.NoError(t, err)
	require.Equal(t, insertPuzzle, getPuzzle)
}

// TestGetPuzzleByArrayStr inserts a random puzzle into the db and attempts to query it by its ArrayStr.
// The test fails if no puzzle was returned or it does not match the inserted puzzle.
func TestGetPuzzleByArrayStr(t *testing.T) {
	insertPuzzle := createRandomPuzzle(t)
	getPuzzle, err := testQueries.GetPuzzleByArrayStr(context.Background(), insertPuzzle.ArrayStr)

	require.NoError(t, err)
	require.Equal(t, insertPuzzle, getPuzzle)
}
