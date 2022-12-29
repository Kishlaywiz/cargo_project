package config

import (
	"backend/service"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	url := "postgres://postgres:postgres@localhost:5432/cargo?sslmode=disable"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("success connection")
	service.InitiateDB(db)
	return db
}
