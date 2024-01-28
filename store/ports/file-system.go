package ports

import "context"

type FileSystem interface {
	RetriveFile(ctx context.Context, nameFiles string) (*[]byte, error)

	SaveFile(ctx context.Context, nameFile string, fileBuff *[]byte) error

	CreateAndZipFiles(ctx context.Context, files map[string]*[]byte, zipFilename string) (*[]byte, error)
}
