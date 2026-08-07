package models

import (
	"github.com/google/uuid"
)

type UserService struct {
	ID        uint `gorm:"primary_key"`
	UserID    uuid.UUID
	ServiceID uint
	StartDate string
	ExpDate   string
	Service   Service `gorm:"foreignKey:ServiceID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
}
