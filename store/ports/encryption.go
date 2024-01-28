package ports

import "context"

type Encryption interface {
	Encrypt(ctx context.Context, file *[]byte) (*[]byte, error)
	Decrypt(ctx context.Context, file *[]byte) (*[]byte, error)
	CalculateFileHash(ctx context.Context, fileContent *[]byte) (string, error)
}