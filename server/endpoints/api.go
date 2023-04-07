package api

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/parhamrou/URL-Shortener/link"
)

func SaveUrl(c echo.Context) error {
	url := c.Param("url")
	if url == "" {
		return c.String(http.StatusBadRequest, "You passed an empty string!")
	}
	link := link.GetShortLink(url)
	err := db.AddLink(link)
	if err != nil {
		return c.String(http.StatusInternalServerError, "There is an internal server problem!")
	}
	return c.String(http.StatusOK, "Your short link is: " + link.ShortLink) 
}


func Redirect(c echo.Context) error {
	shortUrl := c.Param("url")
	if shortUrl == "" {
		return c.String(http.StatusBadRequest, "Your passed url is empty!")
	}
	link, err := db.getLink(shortUrl)
	if err != nil {
		return err
	}
	if link == nil {
		return c.String(http.StatusNotFound, "Invalied URL!")
	} else{
		return c.Redirect(http.StatusOK, link.Original)
	} 
}

