package domain

import (
	"time"
)

type User struct {
	ID        uint `user_id:`
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}
