package main

import (
	"github.com/labstack/echo/v4"
	api "github.com/parhamrou/URL-Shortener/server/endpoints"
	db  "github.com/parhamrou/URL-Shortener/server/database"
)

func main() {
	if db.Connect() != nil {
		return  
	}
	e := echo.New()
	e.POST("/api/shortener", api.SaveUrl)
	e.GET("/api/shortener", api.Redirect)
	e.Logger.Print("Listening on port 8080...\n")
	e.Logger.Fatal(e.Start(":8080"))
}