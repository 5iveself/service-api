package models

type Service struct {
	ID          uint          `json:"id" gorm:"primary_key"`
	ServiceName string        `json:"title"`
	Users       []UserService `gorm:"foreignKey:ServiceIDRef;references:ID"`
}
