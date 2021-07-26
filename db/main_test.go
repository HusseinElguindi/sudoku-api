package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	_ "github.com/lib/pq"
)

// Global test variables
var testDB *sql.DB
var testQueries *Queries

// TestMain prepares and initializes the global test variables.
func TestMain(m *testing.M) {
	var err error

	// TODO: read data from config files
	// Establish a connection with the SQL db
	testDB, err = sql.Open("postgres", "postgres://POSTGRES_USER:POSTGRES_PASSWORD@localhost:5432/sudoku?sslmode=disable")
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	testQueries = New(testDB)

	// Seed the random data generator
	gofakeit.Seed(0)

	os.Exit(m.Run())
}
