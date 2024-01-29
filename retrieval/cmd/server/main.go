// main.go

package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"retrieval/adaptor/repository"
	"retrieval/pkg/envs"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	// for test and develop
	envs.Setup()
}

var (
	KeySecret = []byte(os.Getenv("KEY_ENCRYPTION"))
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(session.Middleware(sessions.NewCookieStore(KeySecret)))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/createuser", createuser)
	e.POST("/login", loginHandler)
	e.GET("/logout", logoutHandler)
	e.GET("/", homeHandler)

	e.POST("/upload", upload)
	e.GET("/retrieve", retrieve)

	// Start server

	e.Logger.Fatal(e.Start(os.Getenv("POSTGRES_GATE_WAY_HOST") + ":" + os.Getenv("SERVER_API_Gate_WAY_PORT")))
}

func loginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	isCorrect, err := repository.CheckPassword(username, password)
	if err != nil || !isCorrect {
		return echo.ErrUnauthorized
	}

	// Create a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires in 1 hour

	// Sign the token with the secret key
	tokenString, err := token.SignedString(KeySecret)
	if err != nil {
		return err
	}

	// Set the token as a cookie
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 1)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "Login successful"})
}

func createuser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	err := repository.CreateUser(username, password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "create user successful"})
}

func logoutHandler(c echo.Context) error {
	// Expire the token cookie
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Expires = time.Now().Add(-time.Hour)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "Logout successful"})
}

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Home Page ")
}

func upload(c echo.Context) error {

	tokenCookie, err := c.Cookie("token")
	if err != nil {
		return echo.ErrUnauthorized
	}

	token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
		return KeySecret, nil
	})
	if err != nil || !token.Valid {
		return echo.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.ErrUnauthorized
	}

	_, ok = claims["username"].(string)
	if !ok {
		return echo.ErrUnauthorized
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	for fieldName, values := range form.File {
		for _, fileHeader := range values {
			file, err := fileHeader.Open()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			defer file.Close()

			part, err := writer.CreateFormFile(fieldName, fileHeader.Filename)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

			if _, err := io.Copy(part, file); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
		}
	}

	for fieldName, values := range form.Value {
		for _, value := range values {
			_ = writer.WriteField(fieldName, value)
		}
	}

	if err := writer.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	resp, err := http.Post("http://"+os.Getenv("SERVER_API_HOST")+":"+os.Getenv("SERVER_API_PORT")+"/upload", writer.FormDataContentType(), &requestBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()
	result := make([]byte, 0)
	resp.Body.Read(result)

	return c.JSONBlob(resp.StatusCode, result)

}

func retrieve(c echo.Context) error {

	tokenCookie, err := c.Cookie("token")
	if err != nil {
		return echo.ErrUnauthorized
	}

	token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
		return KeySecret, nil
	})
	if err != nil || !token.Valid {
		return echo.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.ErrUnauthorized
	}

	_, ok = claims["username"].(string)
	if !ok {
		return echo.ErrUnauthorized
	}

	queryParams := c.QueryParams()

	serviceURL := buildServiceURL("/retrieve", queryParams)

	resp, err := http.Get(serviceURL)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSONBlob(resp.StatusCode, respBody)

}

func buildServiceURL(path string, queryParams url.Values) string {
	u, _ := url.Parse("http://" + os.Getenv("SERVER_API_HOST") + ":" + os.Getenv("SERVER_API_PORT"))

	u.Path = path
	u.RawQuery = queryParams.Encode()
	return u.String()
}
