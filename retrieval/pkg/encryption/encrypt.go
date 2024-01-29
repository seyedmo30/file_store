package encryption

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func GetHash(password string) (string, error) {

	secretKey := os.Getenv("KEY_ENCRYPTION")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+secretKey), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return string(hashedPassword), nil

}

func CheckPassword(inputPassword, hashedPassword string) bool {
	secretKey := os.Getenv("KEY_ENCRYPTION")

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword+string(secretKey)))
	if err == nil {
		return true
	} else {
		return false

	}
}
