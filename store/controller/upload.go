package controller

import (
	"context"
	"store/dto"
	"store/entity"
	"store/ports"
)

type StoreController struct {
	Encryption ports.Encryption
	FileSystem ports.FileSystem
	StorageDB  ports.StorageDB
}

func NewStoreController(encryption ports.Encryption, fileSystem ports.FileSystem, storageDB ports.StorageDB) StoreController {

	return StoreController{Encryption: encryption, FileSystem: fileSystem, StorageDB: storageDB}
}

func (u StoreController) Upload(ctx context.Context, metadata dto.CreateMetadataHttpRequest, file *[]byte, fileName string) (string, int) {

	fileEncrypt, err := u.Encryption.Encrypt(ctx, file)
	if err != nil {
		return "service cant Encrypt file :" + err.Error(), 500
	}
	calculateFileHash, _ := u.Encryption.CalculateFileHash(ctx, file)

	err = u.FileSystem.SaveFile(ctx, calculateFileHash, fileEncrypt)

	if err != nil {
		return "service cant save file :" + err.Error(), 500
	}

	fileFormat := "unknown"
	lastDotIndex := len(fileName) - 1
	for i := lastDotIndex; i >= 0; i-- {
		if fileName[i] == '.' {
			fileFormat = fileName[i+1:]
			break
		}
	}

	createStoreRequest := dto.CreateStoreRequest{entity.Store{FileName: fileName, Type: fileFormat, Name: metadata.Name, Tags: metadata.Tags, Hash: calculateFileHash}}

	err = u.StorageDB.CreateStore(ctx, createStoreRequest)

	if err != nil {
		return "service cant insert to postgres : " + err.Error(), 500
	}

	return "save file successfully .", 201

}

func (u StoreController) Retrieve(ctx context.Context, query dto.RetrieveHttpRequest) (*[]byte, string, int) {
	retrieveStore, err := u.StorageDB.RetrieveStore(ctx, dto.RetrieveStoreRequest{Name: query.Name, Tag: query.Tags})

	if err != nil {
		return nil, "service cant Retrieve database :  " + err.Error(), 500
	}
	mapFiles := make(map[string]*[]byte)

	for _, file := range retrieveStore.Files {

		retriveFile, err := u.FileSystem.RetriveFile(ctx, file.Hash)

		if err != nil {
			return nil, "service cant Retrieve file from system :  " + err.Error(), 500

		}
		fileDecrypt, err := u.Encryption.Decrypt(ctx, retriveFile)

		if err != nil {
			return nil, "service cant Decrypt :  " + err.Error(), 500

		}
		mapFiles[file.FileName] = fileDecrypt

	}

	zipFile, err := u.FileSystem.CreateAndZipFiles(ctx, mapFiles, "response.zip")

	if err != nil {
		return nil, "service zip file :  " + err.Error(), 500

	}

	return zipFile, "successfully", 200

}
