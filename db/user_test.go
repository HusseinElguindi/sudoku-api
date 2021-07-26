package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

// createRandomUser generates a random user and inserts it into the testQueries database.
// The inserted user is returned and guaranteed to be valid. Otherwise, the test (t) is failed.
func createRandomUser(t *testing.T) User {
	// Generate random user data
	params := CreateUserParams{
		FirstName:    gofakeit.FirstName(),
		LastName:     gofakeit.LastName(),
		Username:     gofakeit.Username(),
		PasswordHash: gofakeit.LetterN(25),
	}
	// Insert user into db
	user, err := testQueries.CreateUser(context.Background(), params)

	// Validate the inserted values
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)

	// Compare the inserted values with the original values
	require.Equal(t, user.FirstName, params.FirstName)
	require.Equal(t, user.LastName, params.LastName)
	require.Equal(t, user.Username, params.Username)
	require.Equal(t, user.PasswordHash, params.PasswordHash)
	require.WithinDuration(t, user.CreatedAt, time.Now(), time.Second)

	return user
}

// TestCreateUser inserts a random user into the global test db, failing the test on any errors
// or unexpected results.
func TestCreateUser(t *testing.T) { createRandomUser(t) }

// testCompareUser calls getUserFunc and tests its values for validity, then compares it with
// insertUser for equality.
func testCompareUser(t *testing.T, insertUser User, getUserFunc func() (User, error)) {
	getUser, err := getUserFunc()
	require.NoError(t, err)
	require.Equal(t, insertUser, getUser)
}

// TestGetUserByID inserts a random user into the db and attempts to query it by its ID.
// The test fails if no use was returned or does not match the inserted user.
func TestGetUserByID(t *testing.T) {
	insertUser := createRandomUser(t)
	testCompareUser(t, insertUser, func() (User, error) {
		return testQueries.GetUserByID(context.Background(), insertUser.ID)
	})
}

// TestGetUserByUsername inserts a random user into the db and attempts to query it by its username.
// The test fails if no use was returned or does not match the inserted user.
func GetUserByUsername(t *testing.T) {
	insertUser := createRandomUser(t)
	testCompareUser(t, insertUser, func() (User, error) {
		return testQueries.GetUserByUsername(context.Background(), insertUser.Username)
	})
}

// TestDeleteUserByID inserts and deletes a user by ID and ensures that it was successfully
// removed from the db.
func TestDeleteUserByID(t *testing.T) {
	// Insert a random user into the db
	insertUser := createRandomUser(t)
	// Test querying the inserted user from the db
	testCompareUser(t, insertUser, func() (User, error) {
		return testQueries.GetUserByID(context.Background(), insertUser.ID)
	})

	// Delete the user
	err := testQueries.DeleteUserByID(context.Background(), insertUser.ID)
	require.NoError(t, err)

	// Test querying the inserted user from the db, but expect it to be not found
	user, err := testQueries.GetUserByID(context.Background(), insertUser.ID)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, user)
}

// TestDeleteUserByUsername inserts and deletes a user by username and ensures that it was successfully
// removed from the db.
func TestDeleteUserByUsername(t *testing.T) {
	// Insert a random user into the db
	insertUser := createRandomUser(t)
	// Test querying the inserted user from the db
	testCompareUser(t, insertUser, func() (User, error) {
		return testQueries.GetUserByUsername(context.Background(), insertUser.Username)
	})

	// Delete the user
	err := testQueries.DeleteUserByUsername(context.Background(), insertUser.Username)
	require.NoError(t, err)

	// Test querying the inserted user from the db, but expect it to be not found
	user, err := testQueries.GetUserByUsername(context.Background(), insertUser.Username)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, user)
}
