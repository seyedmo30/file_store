package dto

import "store/entity"

type RetrieveStoreRequest struct {
	Name *string  `json:"name,omitempty"`
	Tag  []string `json:"tag,omitempty"`
}

type RetrieveStoreResponse struct {
	Files []entity.Store `json:"files"`
}

type CreateStoreRequest struct {
	entity.Store
}
