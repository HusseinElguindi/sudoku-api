package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomPuzzle(t *testing.T) Puzzle {
	arrayStr := gofakeit.LetterN(25)
	puzzle, err := testQueries.CreatePuzzle(context.Background(), arrayStr)

	require.NoError(t, err)
	require.NotEmpty(t, puzzle)

	require.NotZero(t, puzzle.ID)
	require.Equal(t, puzzle.ArrayStr, arrayStr)
	require.WithinDuration(t, puzzle.CreatedAt, time.Now(), time.Second)

	return puzzle
}

func TestCreatePuzzle(t *testing.T) {
	createRandomPuzzle(t)
}

func TestGetPuzzleByID(t *testing.T) {
	insertPuzzle := createRandomPuzzle(t)
	getPuzzle, err := testQueries.GetPuzzleByID(context.Background(), insertPuzzle.ID)

	require.NoError(t, err)
	require.Equal(t, insertPuzzle, getPuzzle)
}

func TestGetPuzzleByArrayStr(t *testing.T) {
	insertPuzzle := createRandomPuzzle(t)
	getPuzzle, err := testQueries.GetPuzzleByArrayStr(context.Background(), insertPuzzle.ArrayStr)

	require.NoError(t, err)
	require.Equal(t, insertPuzzle, getPuzzle)
}
