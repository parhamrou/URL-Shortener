package api

import (
	   "net/http"
	   "github.com/labstack/echo/v4"
	   "github.com/parhamrou/URL-Shortener/server/link"
	db "github.com/parhamrou/URL-Shortener/server/database"
)

type body struct {
	URL string `json:"url"`
}


func SaveUrl(c echo.Context) error {
	var reqbody body
	if err := c.Bind(&reqbody); err != nil {
		return err
	}
	url := reqbody.URL
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
	var reqbody body
	if err := c.Bind(&reqbody); err != nil {
		return err
	}
	shortUrl := reqbody.URL
	if shortUrl == "" {
		return c.String(http.StatusBadRequest, "Your passed url is empty!")
	}
	link, err := db.GetLink(shortUrl)
	if err != nil {
		return err
	}
	if link == nil {
		return c.String(http.StatusNotFound, "Invalied URL!")
	} else {
		return c.Redirect(http.StatusTemporaryRedirect, link.OriginalLink)
	} 
}

