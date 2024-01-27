package usecase

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type Encryption struct {
}

func NewEncryption() *Encryption {

	return &Encryption{}
}

func (e Encryption) Encrypt(ctx context.Context, file *[]byte) (*[]byte, error) {

	key := []byte("1234567891234567")

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err

	}

	ciphertext := make([]byte, aes.BlockSize+len(*file))

	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(ciphertext[aes.BlockSize:], *file)

	return &ciphertext, nil
}

func (e Encryption) Decrypt(ctx context.Context, file *[]byte) (*[]byte, error) {

	key := []byte("1234567891234567")

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(*file) < aes.BlockSize {
		panic("Text is too short")
	}

	// Get the 16 byte IV
	iv := (*file)[:aes.BlockSize]

	// Remove the IV from the file
	*file = (*file)[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from file
	stream.XORKeyStream(*file, *file)

	return file, nil

}
