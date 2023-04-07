package main

import (
	"github.com/labstack/echo/v4"
	api "github.com/parhamrou/URL-Shortener/endpoints"
	db  "github.com/parhamrou/URL-Shortener/database"
)

func main() {
	e := echo.New()
	e.POST("/api/shortener", api.SaveUrl)
	e.GET("/api/shortener", api.Redirect)
	e.Logger.Fatal(db.Connect())
	e.Logger.Print("Listening on port 8080...\n")
	e.Logger.Fatal(e.Start(":8080"))
}