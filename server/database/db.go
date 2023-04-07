package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/parhamrou/URL-Shortener/server/link"
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

func AddLink(l *link.Link) error {
	return db.Model(&link.Link{}).Create(l).Error
}


func GetLink(shortLink string) (*link.Link, error) {
	var l link.Link
	result := db.First(&l, shortLink)
	return &l, result.Error
}