package databaseconfig

import (
	"log"
	UserModel "main/cmd/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env files", err)
	}

	dsn := os.Getenv("Database_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error: ", err.Error())
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal("Failed to retrive the db: ", err.Error())
	}
	//checking database connection
	err = sqlDb.Ping()
	if err != nil {
		log.Fatal("Ping error->", err.Error())
	}
	log.Println("Database connection successfull")
	db.AutoMigrate(&UserModel.UserModel{})
	return db
}
