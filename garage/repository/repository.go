package repository

import (
	"section9/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	db.AutoMigrate(&domain.Garage{})
	return &repository{db}
}

func (r *repository) Create(gar *domain.Garage) (*domain.Garage, error) {
	err := r.db.Create(gar).Error
	if err != nil {
		return &domain.Garage{}, err
	}
	return gar, nil
}

func (r *repository) FindAll() ([]domain.Garage, error) {
	var garages []domain.Garage
	err := r.db.Find(&garages).Error
	if err != nil {
		return garages, err
	}
	return garages, nil
}

func (r *repository) Update(garage *domain.Garage) (bool, error) {
	err := r.db.Save(garage).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) FindByID(ID int) (*domain.Garage, error) {
	var garage domain.Garage
	err := r.db.Preload("Cars").Where("id = ?", ID).Find(&garage).Error
	if err != nil {
		return &garage, err
	}
	return &garage, nil
}

func (r *repository) Delete(garage *domain.Garage) (bool, error) {
	if err := r.db.Delete(garage).Error; err != nil {
		return false, err
	}
	return true, nil
}
