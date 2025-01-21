package databaseconfig

import (
	"log"
	UserModel "main/cmd/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	modelToMigrate := []interface{}{
		&UserModel.UserModel{},
		&UserModel.UserDetailModel{},
	}
	log.Println("migrating the database schemas...")
	for _, model := range modelToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("failed to migrate the db %T: %v", model, err)
		}
	}
}
