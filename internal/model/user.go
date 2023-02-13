package model

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	Get(id int) (User, error)
	Create(user *User) error
}

type UserUsecase interface {
	Create(name string) (*User, error)
}
