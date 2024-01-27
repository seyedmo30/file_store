package v1

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
var uploadController controller.UploadController

func init() {
	validate = validator.New()
	encryption = usecase.NewEncryption()
	fileSystem = usecase.NewFileSystem()

	uploadController = controller.NewUploadController(encryption, fileSystem)

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

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error reading the file")
	}
	defer src.Close()

	fileContent, err := io.ReadAll(src)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error reading the file")
	}

	maxSizeFileStr := os.Getenv("max_file_size")
	maxSizeFile, err := strconv.Atoi(maxSizeFileStr)
	if err != nil {

		return c.String(http.StatusInternalServerError, "Error reading the max_file_size")

	}
	if len(fileContent) > int(maxSizeFile) {
		return c.String(http.StatusLengthRequired, "Error reading the max_file_size")

	}

	// controller

	uploadController.Upload(c.Request().Context(), metadata, &fileContent)

	go func() {

	}()
	
	return c.String(http.StatusOK, fmt.Sprintf("File uploaded and encrypted successfully. Metadata: %+v", metadata))
}
