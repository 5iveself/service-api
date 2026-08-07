package models

import (
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	UserID    uuid.UUID
	ServiceID uint
	StartDate time.Time
	ExpDate   time.Time
	User      User    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Service   Service `gorm:"foreignKey:ServiceID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
}
