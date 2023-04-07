# URL-Shortener
A URL shortener written in Go using echo and gorm frameworks to develop restAPI and work with database.
## Requirements
* Postgres
* Go
## Installation
1- Clone the project.

2- Change values in db.go file to match your database configurations.

3- Run `go build` to build the project.

4- Run `./URL-Shortener`
## API
* For shortening a URL, You have to use http POST method with json body with this format:
```
{
  "url": "Your link that you want to be shortened."
}
```
You will get shortened URL in response.
* For going to the original URL, use GET method with json body like above, but with shortened link. You will be redirected to the URL.
