package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
		log.Printf("defaulting to port %s", port)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":" + port))
}
