package service

import (
	"errors"
	"section9/domain"
)

type service struct {
	repository domain.CarRepository
}

func NewService(repository domain.CarRepository) *service {
	return &service{repository}
}

func (s *service) Create(car *domain.Car) (*domain.Car, error) {
	createdCar, err := s.repository.Create(car)
	if err != nil {
		return &domain.Car{}, err
	}
	return createdCar, nil
}

func (s *service) FindAll() ([]domain.Car, error) {
	cars, err := s.repository.FindAll()
	if err != nil {
		return cars, err
	}
	return cars, nil
}

func (s *service) Update(ID int, car *domain.Car) (bool, error) {
	carDB, err := s.repository.FindByID(ID)
	if err != nil {
		return false, err
	}
	if carDB.ID == 0 {
		return false, errors.New("car not found")
	}
	carDB.CarType = car.CarType
	carDB.Label = car.Label
	carDB.Transmision = car.Transmision
	carDB.Fuel = car.Fuel
	carDB.Stock = car.Stock
	carDB.PricePerDay = car.PricePerDay
	carDB.IsActive = car.IsActive
	updatedCar, err := s.repository.Update(carDB)
	if err != nil {
		return false, err
	}
	return updatedCar, nil
}

func (s *service) FindByID(ID int) (*domain.Car, error) {
	car, err := s.repository.FindByID(ID)
	if err != nil {
		return &domain.Car{}, err
	}
	if car.ID == 0 {
		return &domain.Car{}, errors.New("car not found")
	}
	return car, nil
}

func (s *service) Delete(ID int) (bool, error) {
	car, err := s.FindByID(ID)
	if err != nil {
		return false, err
	}
	if car.ID == 0 {
		return false, errors.New("car not found")
	}
	isDeleted, err := s.repository.Delete(car)
	if err != nil {
		return false, err
	}
	return isDeleted, nil
}

func (s *service) UpdateGarage(ID, garageID int) (bool, error) {
	car, err := s.repository.FindByID(ID)
	if err != nil {
		return false, err
	}
	if car.ID == 0 {
		return false, errors.New("car not found")
	}
	updatedCar, err := s.repository.UpdateGarage(ID, garageID)
	if err != nil {
		return false, err
	}
	return updatedCar, nil
}
