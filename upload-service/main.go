package main

import (
	"chapi-backend/chapi-internal/encrypt"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func upload(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	defer src.Close()

	fileName := encrypt.UUIDV1();
	filePath := fmt.Sprintf("images/product/%s%s", fileName, filepath.Ext(file.Filename))
	// Destination
	dst, err := os.Create(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	url := fmt.Sprintf("http://localhost:3002/static/product/%s%s", fileName, filepath.Ext(file.Filename))

	return c.JSON(http.StatusOK, echo.Map{
		"url": url,
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "images/")
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":3002"))
}
