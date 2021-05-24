package main

import (
	// "database/sql"
	// "simpleapi/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")

	})
	e.Logger.Fatal(e.Start(":1323"))
}
