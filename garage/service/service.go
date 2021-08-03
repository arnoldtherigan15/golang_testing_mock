package service

import (
	"errors"
	"section9/domain"
)

type service struct {
	repository domain.GarageRepository
}

func NewService(repository domain.GarageRepository) *service {
	return &service{repository}
}

func (s *service) Create(gar *domain.Garage) (*domain.Garage, error) {
	garage, err := s.repository.Create(gar)
	if err != nil {
		return &domain.Garage{}, err
	}
	return garage, nil
}

func (s *service) FindAll() ([]domain.Garage, error) {
	garages, err := s.repository.FindAll()
	if err != nil {
		return garages, err
	}
	return garages, nil
}

func (s *service) Update(ID int, garage *domain.Garage) (bool, error) {
	garageDB, err := s.repository.FindByID(ID)
	if err != nil {
		return false, err
	}
	if garageDB.ID == 0 {
		return false, errors.New("garage not found")
	}
	garageDB.Owner = garage.Owner
	garageDB.Address = garage.Address
	garageDB.Mobile = garage.Mobile
	updatedGarage, err := s.repository.Update(garageDB)
	if err != nil {
		return false, err
	}
	return updatedGarage, nil
}

func (s *service) FindByID(ID int) (*domain.Garage, error) {
	garage, err := s.repository.FindByID(ID)
	if err != nil {
		return &domain.Garage{}, err
	}
	if garage.ID == 0 {
		return &domain.Garage{}, errors.New("garage not found")
	}
	return garage, nil
}

func (s *service) Delete(ID int) (bool, error) {
	garage, err := s.FindByID(ID)
	if err != nil {
		return false, err
	}
	if garage.ID == 0 {
		return false, errors.New("garage not found")
	}
	isDeleted, err := s.repository.Delete(garage)
	if err != nil {
		return false, err
	}
	return isDeleted, nil
}
