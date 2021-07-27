package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// createRandomUserPuzzle generates a random user and puzzle, creating a userpuzzle between them,
// and inserts it into the testQueries database. The inserted userpuzzle is returned and guaranteed
// to be valid. Otherwise, the test (t) is failed.
func createRandomUserPuzzle(t *testing.T, user User, puzzle Puzzle) UserPuzzle {
	params := CreateUserPuzzleParams{
		UserID:   user.ID,
		PuzzleID: puzzle.ID,
	}
	userPuzzle, err := testQueries.CreateUserPuzzle(context.Background(), params)

	require.NoError(t, err)
	require.Equal(t, userPuzzle.UserID, params.UserID)
	require.Equal(t, userPuzzle.PuzzleID, params.PuzzleID)
	require.WithinDuration(t, userPuzzle.CreatedAt, time.Now(), testTimeThreshold)

	return userPuzzle
}

// TestCreateUserPuzzle inserts a random user, puzzle, and userpuzzle into the global test db,
// failing the test on any errors or unexpected results.
func TestCreateUserPuzzle(t *testing.T) {
	createRandomUserPuzzle(t, createRandomUser(t), createRandomPuzzle(t))
}

// TestGetUserPuzzle inserts a random user, puzzle, and userpuzzle into the global test db,
// then ensures that the inserted userpuzzle is equal to the queried userpuzzle.
func TestGetUserPuzzle(t *testing.T) {
	insertUserPuzzle := createRandomUserPuzzle(t, createRandomUser(t), createRandomPuzzle(t))
	params := GetUserPuzzleParams{
		UserID:   insertUserPuzzle.UserID,
		PuzzleID: insertUserPuzzle.PuzzleID,
	}
	getUserPuzzle, err := testQueries.GetUserPuzzle(context.Background(), params)

	require.NoError(t, err)
	require.Equal(t, insertUserPuzzle, getUserPuzzle)
}

func TestListUserPuzzles(t *testing.T) {
	var (
		n    = 25                  // The number of puzzles to insert and list
		user = createRandomUser(t) // The User of which the UserPuzzles belong
		// Stores the inserted UserPuzzles with PuzzleID as the key
		inserted = make(map[int64]UserPuzzle, n)
	)

	// Create UserPuzzle entries belonging to user
	for i := 0; i < n; i++ {
		userPuzzle := createRandomUserPuzzle(t, user, createRandomPuzzle(t))
		inserted[userPuzzle.PuzzleID] = userPuzzle
	}

	// List UserPuzzles belonging to user
	listed, err := testQueries.ListUserPuzzles(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, listed)

	// Ensure the number of listed and inserted UserPuzzles are the same
	require.Equal(t, len(listed), len(inserted))

	// Iterate throught the listed UserPuzzles and compare them to the inserted UserPuzzles
	for _, up1 := range listed {
		up2, ok := inserted[up1.PuzzleID]
		require.True(t, ok) // Ensure the element is found
		require.NotEmpty(t, up1)
		require.Equal(t, up1, up2)
	}
}

// TestDeleteUserPuzzle inserts and deletes a userpuzzle, ensuring that the userpuzzle was
// removed from the db.
func TestDeleteUserPuzzle(t *testing.T) {
	insertUserPuzzle := createRandomUserPuzzle(t, createRandomUser(t), createRandomPuzzle(t))

	params := GetUserPuzzleParams{
		UserID:   insertUserPuzzle.UserID,
		PuzzleID: insertUserPuzzle.PuzzleID,
	}
	getUserPuzzle, err := testQueries.GetUserPuzzle(context.Background(), params)
	require.NoError(t, err)
	require.Equal(t, insertUserPuzzle, getUserPuzzle)

	err = testQueries.DeleteUserPuzzle(context.Background(), DeleteUserPuzzleParams(params))
	require.NoError(t, err)

	getUserPuzzle, err = testQueries.GetUserPuzzle(context.Background(), params)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, getUserPuzzle)
}
