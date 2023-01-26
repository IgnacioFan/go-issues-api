package model

type User struct {
	ID   uint   `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
