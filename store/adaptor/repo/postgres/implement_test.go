package postgres

import (
	"context"
	"fmt"
	"store/dto"
	"store/entity"
	"store/pkg/envs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	envs.Setup()
}
func TestRetrieveStore(t *testing.T) {

	query := dto.RetrieveStoreRequest{Tag: []string{"111", "Stationery"}}
	results, err := NewPostgres().RetrieveStore(context.Background(), query)
	fmt.Printf("%+v \n", results)
	assert.NoError(t, err)

}

func TestCreateStore(t *testing.T) {
	storeData := entity.Store{
		FileName: "example.txt",
		Name:     "Example Store",
		Tags:     []string{"Electronics", "Gadgets"},
		Type:     "Retail",
		Hash:     "fsdsdfa",
	}
	request := dto.CreateStoreRequest{Store: storeData}

	err := NewPostgres().CreateStore(context.Background(), request)

	assert.NoError(t, err)

}
