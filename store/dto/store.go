package dto

type RetrieveStoreRequest struct {
	Name string   `json:"name"`
	Tag  []string `json:"tag"`
}
