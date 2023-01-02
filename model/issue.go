package model

type Issue struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
}
