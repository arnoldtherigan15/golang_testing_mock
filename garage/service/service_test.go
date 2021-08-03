package service_test

import (
	"errors"
	"section9/domain"
	"section9/domain/mocks"
	_garageService "section9/garage/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGarageService_Create(t *testing.T) {
	mockGarageRepo := new(mocks.GarageRepository)
	mockGarage := &domain.Garage{
		Owner:   "test",
		Address: "california street",
		Mobile:  "08121348584",
	}
	mockEmptyGarage := &domain.Garage{}

	t.Run("success", func(t *testing.T) {
		mockGarageRepo.Mock.On("Create", mock.Anything).Return(mockGarage, nil).Once()
		service := _garageService.NewService(mockGarageRepo)
		garage, err := service.Create(mockGarage)
		assert.NoError(t, err)
		assert.NotNil(t, garage)

		mockGarageRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockGarageRepo.On("Create", mock.Anything).Return(mockEmptyGarage, errors.New("Unexpected")).Once()

		service := _garageService.NewService(mockGarageRepo)

		garage, err := service.Create(mockEmptyGarage)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyGarage, garage)

		mockGarageRepo.AssertExpectations(t)
	})
}

func TestGarageService_FindAll(t *testing.T) {
	mockGarageRepo := new(mocks.GarageRepository)
	mockArrGarage := []domain.Garage{
		domain.Garage{
			Owner:   "test",
			Address: "california street",
			Mobile:  "08121348584",
		},
	}
	mockEmptyGarage := []domain.Garage{
		domain.Garage{},
	}

	t.Run("success", func(t *testing.T) {
		mockGarageRepo.Mock.On("FindAll", mock.Anything).Return(mockArrGarage, nil).Once()
		service := _garageService.NewService(mockGarageRepo)
		garages, err := service.FindAll()
		assert.NoError(t, err)
		assert.NotNil(t, garages)
		mockGarageRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockGarageRepo.On("FindAll", mock.Anything).Return(mockEmptyGarage, errors.New("Unexpected")).Once()

		service := _garageService.NewService(mockGarageRepo)

		garages, err := service.FindAll()

		assert.Error(t, err)
		assert.Equal(t, mockEmptyGarage, garages)

		mockGarageRepo.AssertExpectations(t)
	})
}

func TestGarageService_Update(t *testing.T) {
	mockGarageRepo := new(mocks.GarageRepository)
	mockGarage := &domain.Garage{
		ID:        mock.Anything,
		Owner:     "test",
		Address:   "california street",
		Mobile:    "08121348584",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	// mockEmptyGarage := &domain.Garage{}
	garageID := mock.Anything

	t.Run("success", func(t *testing.T) {
		mockGarageRepo.On("Update", mock.Anything).Return(mockGarage, nil).Once()
		service := _garageService.NewService(mockGarageRepo)
		updatedGarage, err := service.Update(garageID, mockGarage)

		assert.NoError(t, err)
		assert.NotNil(t, updatedGarage)

		mockGarageRepo.AssertExpectations(t)
	})
	// t.Run("error-failed", func(t *testing.T) {
	// 	mockCatRepo.On("UpdateOne", mock.Anything, mock.Anything).Return(mockEmptyCat, errors.New("Unexpected")).Once()

	// 	u := ucase.NewCatUsecase(mockCatRepo, time.Second*2)

	// 	a, err := u.UpdateOne(context.TODO(), mockCat, CatID)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, mockEmptyCat, a)

	// 	mockCatRepo.AssertExpectations(t)
	// })

}
