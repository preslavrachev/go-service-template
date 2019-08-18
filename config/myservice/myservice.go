package myservice

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/preslavrachev/go-service-template/persistence"
	"log"
)

func SetUpDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("DB could not be set up. Reason: %v", err)
	}


	db.AutoMigrate(&persistence.Todo{})

	return db
}
