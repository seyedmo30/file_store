package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

//////////////////////////////////

// type CustomContext struct {
// 	echo.Context
// 	dataStore ports.StorageDB
// }

func NewRouter() *echo.Echo {

	// datastore := postgres.NewPostgres()

	e := echo.New()
	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		cc := &CustomContext{c, datastore}
	// 		return next(cc)
	// 	}
	// })
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// routers

	e.POST("/upload", UploadHandler)

	// e.GET("/main", CrontabMain)

	// Route to serve a media file from a byte buffer
	e.GET("/media", func(c echo.Context) error {
		// Replace `fileContent` with the actual byte buffer containing your file
		fileContent := []byte("Your binary file content goes here.")

		// Set the content type based on your file type (e.g., "video/mp4", "image/jpeg", etc.)
		contentType := "application/octet-stream"

		// Set the response header
		c.Response().Header().Set("Content-Type", contentType)

		// Send the file content as the response
		return c.Blob(http.StatusOK, contentType, fileContent)
	})

	return e
}
