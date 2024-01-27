package usecase

import (
	"bytes"
	"context"
	"encoding/binary"
	"log"
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
		log.Fatal(err)
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
