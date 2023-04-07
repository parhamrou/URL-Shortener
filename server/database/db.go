package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host	 =	"localhost"
	user 	 =  "parhamrou"
	password =	"parham1381"
	dbName 	 =  "url_shortener"
	port 	 =   5432
)

var db *gorm.DB

func Connect() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

func AddLink(link *Link) {
	link
}