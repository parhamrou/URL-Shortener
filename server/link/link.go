package link

import (
	"crypto/md5"
	"encoding/base64"
	"strings"
)

const urlBase = "Parham.me/"

type Link struct {
	OriginalLink   string 
	ShortLink  string 
}

// This function gets original link and returns a Link object
func GetShortLink(url string) *Link {
	trueLink := url
	if !strings.Contains(trueLink, "http://") {
		trueLink = "http://" + trueLink
	} 
	bytes := md5.Sum([]byte(trueLink))
	shortLink := "http://" + urlBase + base64.StdEncoding.EncodeToString(bytes[:])[:6]
	link := Link{
		OriginalLink: trueLink,
		ShortLink: shortLink,
	}
	return &link
}