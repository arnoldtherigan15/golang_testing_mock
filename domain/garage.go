package domain

import (
	"time"

	"gorm.io/gorm"
)

type Garage struct {
	ID        int            `json:"id"`
	Owner     string         `json:"owner" validate:"required"`
	Address   string         `json:"address" validate:"required"`
	Mobile    string         `json:"mobile" validate:"required"`
	Cars      []Car          `gorm:"foreignKey:GarageID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type GarageRepository interface {
	Create(gar *Garage) (*Garage, error)
	Update(gar *Garage) (bool, error)
	FindByID(ID int) (*Garage, error)
	FindAll() ([]Garage, error)
	Delete(gar *Garage) (bool, error)
}

type GarageService interface {
	Create(gar *Garage) (*Garage, error)
	Update(ID int, gar *Garage) (bool, error)
	FindByID(ID int) (*Garage, error)
	FindAll() ([]Garage, error)
	Delete(ID int) (bool, error)
}
