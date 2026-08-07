package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Username string    `json:"username" gorm:"not null"`
	Password string    `json:"password" gorm:"not null"`
}
