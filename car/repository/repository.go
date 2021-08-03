package repository

import (
	"section9/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	db.AutoMigrate(&domain.Car{})
	return &repository{db}
}

func (r *repository) Create(car *domain.Car) (*domain.Car, error) {
	err := r.db.Omit("GarageID").Create(car).Error
	if err != nil {
		return &domain.Car{}, err
	}
	return car, nil
}

func (r *repository) FindAll() ([]domain.Car, error) {
	var cars []domain.Car
	err := r.db.Find(&cars).Error
	if err != nil {
		return cars, err
	}
	return cars, nil
}

func (r *repository) Update(car *domain.Car) (bool, error) {
	err := r.db.Omit("GarageID").Save(car).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) FindByID(ID int) (*domain.Car, error) {
	var car domain.Car
	err := r.db.Preload("Garage").Where("id = ?", ID).Find(&car).Error
	if err != nil {
		return &car, err
	}
	return &car, nil
}

func (r *repository) UpdateGarage(ID, garageID int) (bool, error) {
	err := r.db.Model(&domain.Car{}).Where("id = ?", ID).Update("garage_id", garageID).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) Delete(car *domain.Car) (bool, error) {
	if err := r.db.Delete(car).Error; err != nil {
		return false, err
	}
	return true, nil
}
