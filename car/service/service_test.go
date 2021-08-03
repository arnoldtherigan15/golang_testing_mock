package service_test

import (
	"errors"
	_carService "section9/car/service"
	"section9/domain"
	"section9/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCarService_Create(t *testing.T) {
	mockCarRepo := new(mocks.CarRepository)
	mockCar := &domain.Car{
		CarType:     "car type",
		Label:       "car label",
		Transmision: "car transmission",
		Fuel:        "car fuel",
		Stock:       2,
		PricePerDay: 50000,
		IsActive:    true,
	}
	mockEmptyCar := &domain.Car{}

	t.Run("success", func(t *testing.T) {
		mockCarRepo.Mock.On("Create", mock.Anything).Return(mockCar, nil).Once()
		service := _carService.NewService(mockCarRepo)
		garage, err := service.Create(mockCar)
		assert.NoError(t, err)
		assert.NotNil(t, garage)

		mockCarRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCarRepo.On("Create", mock.Anything).Return(mockEmptyCar, errors.New("Unexpected")).Once()

		service := _carService.NewService(mockCarRepo)

		garage, err := service.Create(mockEmptyCar)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyCar, garage)

		mockCarRepo.AssertExpectations(t)
	})
}

func TestCarService_FindAll(t *testing.T) {
	mockCarRepo := new(mocks.CarRepository)
	mockArrGarage := []domain.Car{
		domain.Car{
			CarType:     "car type",
			Label:       "car label",
			Transmision: "car transmission",
			Fuel:        "car fuel",
			Stock:       2,
			PricePerDay: 50000,
			IsActive:    true,
		},
	}
	mockEmptyCar := []domain.Car{
		domain.Car{},
	}

	t.Run("success", func(t *testing.T) {
		mockCarRepo.Mock.On("FindAll", mock.Anything).Return(mockArrGarage, nil).Once()
		service := _carService.NewService(mockCarRepo)
		garages, err := service.FindAll()
		assert.NoError(t, err)
		assert.NotNil(t, garages)
		mockCarRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCarRepo.On("FindAll", mock.Anything).Return(mockEmptyCar, errors.New("Unexpected")).Once()

		service := _carService.NewService(mockCarRepo)

		garages, err := service.FindAll()

		assert.Error(t, err)
		assert.Equal(t, mockEmptyCar, garages)

		mockCarRepo.AssertExpectations(t)
	})
}

func TestCarService_Update(t *testing.T) {
	mockCarRepo := new(mocks.CarRepository)
	mockCar := &domain.Car{
		ID:          1,
		CarType:     "car type",
		Label:       "car label",
		Transmision: "car transmission",
		Fuel:        "car fuel",
		Stock:       2,
		PricePerDay: 50000,
		IsActive:    true,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	mockEmptyCar := &domain.Car{}

	t.Run("success", func(t *testing.T) {
		mockCarRepo.On("FindByID", mock.Anything).Return(mockCar, nil).Once()
		mockCarRepo.On("Update", mock.Anything, mock.Anything).Return(true, nil).Once()
		service := _carService.NewService(mockCarRepo)
		isUpdated, err := service.Update(1, mockCar)

		assert.NoError(t, err)
		assert.Equal(t, isUpdated, true)
		assert.NotNil(t, isUpdated)

		mockCarRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCarRepo.On("FindByID", mock.Anything).Return(mockCar, nil).Once()
		mockCarRepo.On("Update", mock.Anything).Return(false, errors.New("unexpected")).Once()

		service := _carService.NewService(mockCarRepo)
		isUpdated, err := service.Update(1, mockEmptyCar)
		assert.Error(t, err)
		assert.Equal(t, isUpdated, false)

		mockCarRepo.AssertExpectations(t)
	})

}

func TestCarService_FindByID(t *testing.T) {
	mockCarRepo := new(mocks.CarRepository)
	mockCar := &domain.Car{
		ID:          1,
		CarType:     "car type",
		Label:       "car label",
		Transmision: "car transmission",
		Fuel:        "car fuel",
		Stock:       2,
		PricePerDay: 50000,
		IsActive:    true,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	mockEmptyCar := &domain.Car{}

	t.Run("success", func(t *testing.T) {
		mockCarRepo.On("FindByID", mock.Anything).Return(mockCar, nil).Once()
		service := _carService.NewService(mockCarRepo)
		garage, err := service.FindByID(1)

		assert.NoError(t, err)
		assert.Equal(t, mockCar, garage)
		assert.NotNil(t, garage)

		mockCarRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCarRepo.On("FindByID", mock.Anything).Return(mockEmptyCar, errors.New("unexpected")).Once()

		service := _carService.NewService(mockCarRepo)
		garage, err := service.FindByID(1)
		assert.Error(t, err)
		assert.Equal(t, garage, mockEmptyCar)

		mockCarRepo.AssertExpectations(t)
	})

}

func TestCarService_Delete(t *testing.T) {
	mockCarRepo := new(mocks.CarRepository)
	mockCar := &domain.Car{
		ID:          1,
		CarType:     "car type",
		Label:       "car label",
		Transmision: "car transmission",
		Fuel:        "car fuel",
		Stock:       2,
		PricePerDay: 50000,
		IsActive:    true,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockCarRepo.On("FindByID", mock.Anything).Return(mockCar, nil).Once()
		mockCarRepo.On("Delete", mock.Anything).Return(true, nil).Once()
		service := _carService.NewService(mockCarRepo)
		isDeleted, err := service.Delete(1)

		assert.NoError(t, err)
		assert.Equal(t, isDeleted, true)
		assert.NotNil(t, isDeleted)

		mockCarRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCarRepo.On("FindByID", mock.Anything).Return(mockCar, nil).Once()
		mockCarRepo.On("Delete", mock.Anything).Return(false, errors.New("unexpected")).Once()
		service := _carService.NewService(mockCarRepo)
		isDeleted, err := service.Delete(1)
		assert.Error(t, err)
		assert.Equal(t, isDeleted, false)

		mockCarRepo.AssertExpectations(t)
	})

}

func TestCarService_UpdateGarage(t *testing.T) {
	mockCarRepo := new(mocks.CarRepository)
	mockCar := &domain.Car{
		ID:          1,
		CarType:     "car type",
		Label:       "car label",
		Transmision: "car transmission",
		Fuel:        "car fuel",
		Stock:       2,
		PricePerDay: 50000,
		IsActive:    true,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	// mockEmptyCar := &domain.Car{}

	t.Run("success", func(t *testing.T) {
		mockCarRepo.On("FindByID", mock.Anything).Return(mockCar, nil).Once()
		mockCarRepo.On("UpdateGarage", mock.Anything, mock.Anything).Return(true, nil).Once()
		service := _carService.NewService(mockCarRepo)
		isUpdated, err := service.UpdateGarage(1, 1)

		assert.NoError(t, err)
		assert.Equal(t, isUpdated, true)
		assert.NotNil(t, isUpdated)

		mockCarRepo.AssertExpectations(t)
	})
	// t.Run("error-failed", func(t *testing.T) {
	// 	mockCarRepo.On("FindByID", mock.Anything).Return(mockCar, nil).Once()
	// 	mockCarRepo.On("Update", mock.Anything).Return(false, errors.New("unexpected")).Once()

	// 	service := _carService.NewService(mockCarRepo)
	// 	isUpdated, err := service.Update(1, mockEmptyCar)
	// 	assert.Error(t, err)
	// 	assert.Equal(t, isUpdated, false)

	// 	mockCarRepo.AssertExpectations(t)
	// })

}
