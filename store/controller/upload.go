package controller

import (
	"context"
	"store/dto"
	"store/ports"
)

type UploadController struct {
	Encryption ports.Encryption
	FileSystem ports.FileSystem
}

func NewUploadController(encryption ports.Encryption, fileSystem ports.FileSystem) UploadController {

	return UploadController{Encryption: encryption, FileSystem: fileSystem}
}

func (u UploadController) Upload(ctx context.Context, metadata dto.CreateMetadataHttpRequest, file *[]byte) (string, int) {

	fileEncrypt, err := u.Encryption.Encrypt(ctx, file)
	if err != nil {
		return "service cant Encrypt file ", 500
	}

	err = u.FileSystem.SaveFile(ctx, "salam", fileEncrypt)

	if err != nil {
		return "service cant save file ", 500
	}

	return "service cant Encrypt file ", 500

}
