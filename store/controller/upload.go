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

	createStoreRequest := dto.CreateStoreRequest{Store: entity.Store{FileName: fileName, Type: fileFormat, Name: metadata.Name, Tags: metadata.Tags, Hash: calculateFileHash}}

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
	hashList := make([]string, 0, 1)
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

		hashList = append(hashList, file.Hash)

	}
	if len(hashList) == 0 {
		retrieveStore , err :=u.StorageDB.RetrieveFirstStore(ctx)


		if err != nil {
			return nil, "service cant Retrieve file from system :  " + err.Error(), 500

		}

		retriveFile, err := u.FileSystem.RetriveFile(ctx, retrieveStore.Hash)

		if err != nil {
			return nil, "service cant Retrieve file from system :  " + err.Error(), 500

		}
		
		fileDecrypt, err := u.Encryption.Decrypt(ctx, retriveFile)

		if err != nil {
			return nil, "service cant Decrypt :  " + err.Error(), 500

		}
		mapFiles[retrieveStore.FileName] = fileDecrypt

		hashList = append(hashList, retrieveStore.Hash)
		if len(hashList) == 0{

			return nil, "service cant find file  ", 400
		}

	}
	zipFile, err := u.FileSystem.CreateAndZipFiles(ctx, mapFiles, "response.zip")

	if err != nil {
		return nil, "service cant zip file :  " + err.Error(), 500

	}

	err = u.StorageDB.DeleteStore(ctx, hashList)

	if err != nil {
		return nil, "service  cant delete from database :  " + err.Error(), 500
	}

	err = u.FileSystem.DeleteFiles(ctx, hashList)
	if err != nil {
		return nil, "service  cant delete system files :  " + err.Error(), 500
	}

	return zipFile, "successfully", 200

}
