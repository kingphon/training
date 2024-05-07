package main

import (
	"go-sql/internal/dump"
	dbpg "go-sql/internal/pg/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	dbpg.ConnectDB()
	dump.Init()
}

func main() {
	e := echo.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
