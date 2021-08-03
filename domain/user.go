package domain

import (
	"section9/user/formatter"
	"time"

	"section9/user/input"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id"`
	FullName  string         `json:"fullname" form:"fullname" validate:"required"`
	Mobile    string         `json:"mobile" form:"mobile" validate:"required"`
	Address   string         `json:"address" form:"address" validate:"required"`
	Email     string         `json:"email" form:"email" validate:"required"`
	IDCard    string         `json:"id_card" form:"id_card" validate:"required"`
	Role      string         `json:"role"`
	Password  string         `json:"password" form:"password" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	Update(user *User) (bool, error)
	FindByID(ID int) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll() ([]*User, error)
	Delete(user *User) (bool, error)
}

type UserService interface {
	Create(user *User) (*User, error)
	Update(ID int, user *input.UpdateInput) (bool, error)
	FindByID(ID int) (*User, error)
	FindAll() ([]*User, error)
	Delete(ID int) (bool, error)
	Login(input input.LoginInput) (*formatter.LoginResponse, error)
}
