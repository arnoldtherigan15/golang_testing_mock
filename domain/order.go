package domain

import (
	"time"

	"section9/order/formatter"
	"section9/order/helpers"
	"section9/order/input"

	"gorm.io/gorm"
)

type Order struct {
	ID            int            `json:"id"`
	EstimatedDays int            `json:"estimated_days"`
	WithDriver    bool           `json:"with_driver"`
	Status        string         `json:"status"`
	CarID         int            `json:"car_id"`
	Car           Car            `gorm:"foreignKey:CarID" json:"car"`
	UserID        int            `json:"user_id"`
	User          User           `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

type OrderRepository interface {
	Create(order *Order) (*Order, error)
	Done(order *Order) (bool, error)
	FindByID(ID int) (*Order, error)
	FindAll(pagination helpers.Pagination) (*helpers.Pagination, error)
	// Delete(order *Order) (bool, error)
}

type OrderService interface {
	Create(order *input.Order) (*formatter.Formatter, error)
	Done(ID *input.DoneInput) (*formatter.DoneFormatter, error)
	FindAll(pagination helpers.Pagination) (*helpers.Pagination, error)
	// FindByID(ID int) (*Order, error)
	// FindAll() (*[]Order, error)
	// Delete(ID int) (bool, error)
}
