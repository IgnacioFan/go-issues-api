package model

type Issue struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      User   `gorm:"foreignKey:ID" json:"author"`
}
