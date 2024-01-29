package encryption

import (
	"os"
	"testing"
)

func TestGetHashAndCheckPassword(t *testing.T) {
	// Set up a temporary environment variable for testing
	os.Setenv("KEY_ENCRYPTION", "test_key")

	// Test GetHash function
	password := "test_password"
	hashedPassword, err := GetHash(password)

	if err != nil {
		t.Fatalf("GetHash function failed: %v", err)
	}

	// Test GetPassword function
	match := CheckPassword(password, hashedPassword)
	if !match {
		t.Error("CheckPassword function failed: passwords do not match")
	}

	// Test case where password doesn't match
	nonMatchingPassword := "wrong_password"
	nonMatching := CheckPassword(nonMatchingPassword, hashedPassword)
	if nonMatching {
		t.Error("CheckPassword function failed: non-matching passwords considered as matching")
	}

	// Clean up environment variable
	os.Unsetenv("KEY_ENCRYPTION")
}
