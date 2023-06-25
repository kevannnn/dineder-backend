package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = InitDB()

func InitDB() *gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=a12b34c56 dbname=dineders port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
