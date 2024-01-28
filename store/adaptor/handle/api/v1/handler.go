package v1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"store/adaptor/repo/postgres"
	"store/controller"
	"store/dto"
	"store/usecase"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate
var encryption *usecase.Encryption
var fileSystem *usecase.FileSystem
var postgresSetup postgres.Setup
var storeController controller.StoreController

func init() {
	validate = validator.New()
	encryption = usecase.NewEncryption()
	fileSystem = usecase.NewFileSystem()
	postgresSetup = postgres.NewPostgres()
	storeController = controller.NewStoreController(encryption, fileSystem, postgresSetup)

}

func UploadHandler(c echo.Context) error {
	// bind and validate
	var metadata dto.CreateMetadataHttpRequest
	if err := c.Bind(&metadata); err != nil {
		return c.String(http.StatusBadRequest, "Error binding metadata")
	}

	if err := validate.Struct(metadata); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error()))
	}
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "Error retrieving the file")
	}

	fileName := file.Filename
	maxSizeFileNameStr := os.Getenv("MAX_FILE_NAME_SIZE")
	maxSizeFileName, err := strconv.Atoi(maxSizeFileNameStr)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error reading the max_file_name_size")
	}

	if len(fileName) > int(maxSizeFileName) {
		return c.String(http.StatusLengthRequired, "file name is too long")
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error reading the file")
	}
	defer src.Close()

	fileContent, err := io.ReadAll(src)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error reading the file")
	}

	maxSizeFileStr := os.Getenv("MAX_FILE_SIZE")
	maxSizeFile, err := strconv.Atoi(maxSizeFileStr)
	if err != nil {

		return c.String(http.StatusInternalServerError, "Error reading the max_file_size")

	}
	if len(fileContent) > int(maxSizeFile) {
		return c.String(http.StatusLengthRequired, "file size is too long")

	}

	// controller

	message, statusCode := storeController.Upload(c.Request().Context(), metadata, &fileContent, fileName)

	go func() {

	}()

	return c.String(statusCode, message)
}

func RetrieveHandler(c echo.Context) error {
	// bind
	req := new(dto.RetrieveHttpRequest)

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "Error binding query parameters")
	}
	// controller
	fileContent, message, statusCode := storeController.Retrieve(c.Request().Context(), *req)
	if statusCode != 200 {

		return c.String(statusCode, message)
	}
	// Set the content type based on your file type (e.g., "video/mp4", "image/jpeg", etc.)
	contentType := "application/octet-stream"

	// Set the response header
	c.Response().Header().Set("Content-Type", contentType)

	// Send the file content as the response
	return c.Blob(http.StatusOK, contentType, *fileContent)

}
