package database

import (
	"log"
	"main/cmd/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=9843 dbname=springboot port=5432 sslmode=disable"
	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error: ", err.Error())
	}
	err = DB.AutoMigrate(models.UserModel{})
	if err != nil {
		log.Fatal("Migration error:", err.Error())
	}

	sqlDb, err := DB.DB()

	if err != nil {
		log.Fatal(err)
	}
	if err = sqlDb.Ping(); err != nil {
		log.Fatal("Database ping error:", err.Error())
	} else {
		log.Println("Database connected to the port 5432")
	}
	return DB
}
