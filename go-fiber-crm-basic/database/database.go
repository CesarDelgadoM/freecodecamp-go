package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("connection to sqlite failed:", err)
	}
	DB = db
	log.Println("connection succesful to sqlite")
}

func AutoMigrate(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatal("migration failed:", err)
	}
	log.Println("migration succesful")
}
