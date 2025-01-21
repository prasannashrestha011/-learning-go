package UserDTOS

import "time"

type CreateUserDTO struct {
	UserId      *string        `json:"user_id"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Email       string         `json:"email"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	UserDetails UserDetailsDTO `json:"user_details"`
}
type UserDetailsDTO struct {
	DetailId      *string `json:"detail_id"`
	Address       string  `json:"address"`
	ContactNumber int64   `json:"contact_number"`
}
