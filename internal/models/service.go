package models

type Service struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	ServiceName string `json:"title"`
	Price       int
}
