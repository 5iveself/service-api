package models

import "time"

type UserService struct {
	UserID    uint
	ServiceID uint
	StartDate time.Time
	ExpDate   time.Time

	UserIDRef    User    `gorm:"foreignKey:Username;references:Username;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ServiceIDRef Service `gorm:"foreignKey:ServiceName;references:ServiceName;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
