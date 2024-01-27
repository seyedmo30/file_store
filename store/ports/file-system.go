package ports

import "context"

type FileSystem interface {
	RetriveFile(ctx context.Context, nameFile string) (*[]byte, error)

	SaveFile(ctx context.Context, nameFile string, fileBuff *[]byte) error
}
