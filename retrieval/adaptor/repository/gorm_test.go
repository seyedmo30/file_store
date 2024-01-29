package repository

// func TestGorm(t *testing.T) {
// 	Gorm()
// }

import (
	"retrieval/entity"
	"retrieval/pkg/envs"
	"testing"
)

func TestCreateUserAndCheckPassword(t *testing.T) {
	envs.Setup()
	// Ensure your database is set up for testing, you may use a test database or mock the database.
	// Initialize the database connection and migrate the necessary tables.
	db, err := GetDBInstance()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Test CreateUser function
	username := "admin"
	password := "admin"
	err = CreateUser(username, password)
	if err != nil {
		t.Fatalf("CreateUser function failed: %v", err)
	}
	// Fetch the user from the database and check if it was created correctly
	var fetchedUser entity.User
	db.First(&fetchedUser, "username = ?", username)
	if fetchedUser.Username != username {
		t.Errorf("CreateUser function failed: user not found in the database")
	}

	// Test CheckPassword function
	correctPassword := "testpassword"
	incorrectPassword := "wrongpassword"

	// Check correct password
	match, err := CheckPassword(username, correctPassword)
	if err != nil {
		t.Fatalf("CheckPassword function failed: %v", err)
	}
	if !match {
		t.Error("CheckPassword function failed: incorrect result for correct password")
	}

	// Check incorrect password
	match, err = CheckPassword(username, incorrectPassword)
	if err != nil {
		t.Fatalf("CheckPassword function failed: %v", err)
	}
	if match {
		t.Error("CheckPassword function failed: incorrect result for incorrect password")
	}
}
