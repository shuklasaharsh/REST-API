package database

import (
	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	"log"
)

var db *gorm.DB

func ConnectDB(dbUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	green:= color.New(color.BgGreen)
	green.Println("DB Connected Successfully")

	return db, err
}

func GetDB() *gorm.DB {
	return db
}