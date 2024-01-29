package dto

type CreateMetadataHttpRequest struct {
	Name string   `form:"name" validate:"required"`
	Tags []string `form:"tags" `
	Type string   `form:"type" validate:"required"`
}

type RetrieveHttpRequest struct {
	Name *string   `query:"name" `
	Tags []string `query:"tags" `
}
