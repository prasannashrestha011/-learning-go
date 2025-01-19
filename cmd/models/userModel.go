package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	UserId    string    `gorm:"primarykey" json:"user_id"`
	Username  string    `gorm:"not null;size:100" json:"username"`
	Email     string    `gorm:"unique;notnull;size:100" json:"email"`
	Password  string    `gorm:"size:250" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (user *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("generating userid......")
	user.UserId = uuid.New().String()[:16]
	return nil
}
