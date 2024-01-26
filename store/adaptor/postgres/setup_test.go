package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDB(t *testing.T) {
	var err error
	// Assuming that there is a PostgreSQL server running with the test user and database
	db := GetDB()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Check if the connection is valid
	err = db.Ping()
	assert.NoError(t, err)

	// Perform additional tests or queries as needed
}

// func TestInitializeDB(t *testing.T) {
// 	// Create a mock SQL database for testing purposes
// 	mockDB, err := sql.Open("sqlmock")
// 	assert.NoError(t, err)

// 	// Set the mockDB for testing
// 	db = mockDB

// 	// Call initializeDB
// 	initializeDB()

// 	// Check if db is assigned correctly
// 	assert.Equal(t, mockDB, db)
// }
