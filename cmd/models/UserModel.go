package UserModel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	UserId      string           `gorm:"primarykey" json:"user_id"`
	Username    string           `gorm:"unique;size:100;" json:"username"`
	Password    string           `gorm:"size:100" json:"password"`
	Email       string           `gorm:"unique;size:255" json:"email"`
	CreatedAt   time.Time        `gorm:"createdAt" json:"created_at"`
	UpdatedAt   time.Time        `gorm:"updatedAt" json:"updated_at"`
	UserDetails *UserDetailModel `gorm:"foreignKey:DetailId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_details"`
}
type UserDetailModel struct {
	DetailId      string `gorm:"primarykey" json:"detail_id"`
	Address       string `gorm:"size:255;" json:"address"`
	ContactNumber int64  `json:"contact_number"`
}

func (user *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	user.UserId = uuid.New().String()[:16]
	return nil
}
