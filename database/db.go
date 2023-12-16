package database

import (
	"Assignment-3/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost     = "localhost"
	dbPort     = 5434
	dbUser     = "postgres"
	dbPassword = "23032002"
	dbName     = "Assignment-3"
	db         *gorm.DB
)

func StartDB() {
	Config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error

	db, err = gorm.Open(postgres.Open(Config), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.SensorData{})
	fmt.Println("Connected to the database")
}

func GetDB() *gorm.DB {
	return db
}
