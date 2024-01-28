package usecase

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/binary"
	"os"
	"store/pkg/logs"
)

type FileSystem struct {
}

func NewFileSystem() *FileSystem {

	return &FileSystem{}
}

func (f FileSystem) RetriveFile(ctx context.Context, nameFile string) (*[]byte, error) {

	content, err := os.ReadFile(nameFile)
	if err != nil {
		return nil, err
	}
	return &content, nil

}
func writeNextBytes(file *os.File, bytes []byte) {

	_, err := file.Write(bytes)

	if err != nil {
		logs.Connect().Error(err.Error())
	}

}

func (f FileSystem) SaveFile(ctx context.Context, nameFile string, fileBuff *[]byte) error {
	file, err := os.Create(nameFile)
	if err != nil {
		logs.Connect().Error(err.Error())
		return err

	}
	defer file.Close()

	var bin_buf bytes.Buffer
	binary.Write(&bin_buf, binary.BigEndian, *fileBuff)

	writeNextBytes(file, bin_buf.Bytes())

	return nil
}

func (f FileSystem) CreateAndZipFiles(ctx context.Context, files map[string]*[]byte, zipFilename string) (*[]byte, error) {
	var zipBuffer bytes.Buffer

	zipWriter := zip.NewWriter(&zipBuffer)

	for filename, contentPtr := range files {
		fileWriter, err := zipWriter.Create(filename)
		if err != nil {
			return nil, err
		}

		_, err = fileWriter.Write(*contentPtr)
		if err != nil {
			return nil, err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return nil, err
	}
	result := zipBuffer.Bytes()
	return &result, nil
}

func (f FileSystem) DeleteFiles(ctx context.Context, fileNames []string) error {
	var err error
	for _, fileName := range fileNames {
		err = os.Remove(fileName)

	}
	return err
}
