package model

type User struct {
	ID   uint   `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type UserRepository interface {
	Get(id int) (User, error)
	Create(user *User) error
}
