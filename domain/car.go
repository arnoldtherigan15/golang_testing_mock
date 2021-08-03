package domain

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	ID          int            `json:"id"`
	CarType     string         `json:"car_type" form:"car_type" validate:"required"`
	Label       string         `json:"label" form:"label" validate:"required"`
	Transmision string         `json:"transmision" form:"transmision" validate:"required"`
	Fuel        string         `json:"fuel" form:"fuel" validate:"required"`
	Stock       int            `json:"stock" form:"stock"`
	PricePerDay int            `json:"price_per_day" form:"price_per_day" validate:"required"`
	IsActive    bool           `json:"is_active" form:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	GarageID    int            `json:"garage_id"`
	Garage      *Garage        `gorm:"foreignKey:GarageID;" json:"garage"`
	// Orders      []Order `gorm:"foreignKey:CarID"`
}

type CarRepository interface {
	Create(car *Car) (*Car, error)
	UpdateGarage(ID, garageID int) (bool, error)
	Update(car *Car) (bool, error)
	FindByID(ID int) (*Car, error)
	FindAll() ([]Car, error)
	Delete(car *Car) (bool, error)
}

type CarService interface {
	Create(car *Car) (*Car, error)
	UpdateGarage(ID, garageID int) (bool, error)
	Update(ID int, car *Car) (bool, error)
	FindByID(ID int) (*Car, error)
	FindAll() ([]Car, error)
	Delete(ID int) (bool, error)
}
