package ports

import (
	"context"
	"store/dto"
)

type StorageDB interface {
	CreateStore(ctx context.Context, request dto.CreateStoreRequest) error

	RetrieveStore(ctx context.Context, query dto.RetrieveStoreRequest) (dto.RetrieveStoreResponse, error)

	DeleteStore(ctx context.Context, hashList []string) error
}
