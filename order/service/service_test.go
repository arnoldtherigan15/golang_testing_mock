package service_test

import (
	"errors"
	"section9/domain"
	"section9/domain/mocks"
	"section9/order/helpers"
	"section9/order/input"
	_orderService "section9/order/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrderService_Create(t *testing.T) {
	mockOrderRepo := new(mocks.OrderRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockCarepo := new(mocks.CarRepository)

	mockOrder := &domain.Order{
		EstimatedDays: 3,
		WithDriver:    true,
		Status:        "pending",
		CarID:         1,
		UserID:        1,
	}
	mockOrderInput := &input.Order{
		EstimatedDays: 3,
		WithDriver:    true,
		CarID:         1,
		UserID:        1,
	}
	mockEmptyOrder := &domain.Order{}
	mockEmptyInputOrder := &input.Order{}
	mockUser := &domain.User{
		ID:       1,
		FullName: "arnold",
		Address:  "california street",
		Mobile:   "08121348584",
		Email:    "arnold@mail.com",
		IDCard:   "id card",
		Role:     "employee",
	}
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
		mockUserRepo.Mock.On("FindByID", mock.Anything).Return(mockUser, nil).Once()
		mockCarepo.Mock.On("FindByID", mock.Anything).Return(mockCar, nil).Once()
		mockCarepo.Mock.On("Update", mock.Anything).Return(true, nil).Once()
		mockOrderRepo.Mock.On("Create", mock.Anything).Return(mockOrder, nil).Once()

		service := _orderService.NewService(mockOrderRepo, mockCarepo, mockUserRepo)
		order, err := service.Create(mockOrderInput)
		assert.NoError(t, err)
		assert.NotNil(t, order)

		mockOrderRepo.AssertExpectations(t)
		mockCarepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.Mock.On("FindByID", mock.Anything).Return(mockUser, nil).Once()
		mockCarepo.Mock.On("FindByID", mock.Anything).Return(mockCar, nil).Once()
		mockOrderRepo.Mock.On("Create", mock.Anything).Return(mockEmptyOrder, errors.New("unexpected")).Once()

		service := _orderService.NewService(mockOrderRepo, mockCarepo, mockUserRepo)

		_, err := service.Create(mockEmptyInputOrder)

		assert.Error(t, err)

		mockOrderRepo.AssertExpectations(t)
		mockCarepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})
}
func TestOrderService_Done(t *testing.T) {
	mockOrderRepo := new(mocks.OrderRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockCarepo := new(mocks.CarRepository)
	mockOrder := &domain.Order{
		ID:            1,
		EstimatedDays: 3,
		WithDriver:    true,
		Status:        "pending",
		CarID:         1,
		UserID:        1,
		UpdatedAt:     time.Now(),
		CreatedAt:     time.Now(),
	}
	mockOrderInput := &input.DoneInput{
		OrderID: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockOrderRepo.Mock.On("FindByID", mock.Anything).Return(mockOrder, nil).Once()
		mockOrderRepo.Mock.On("Done", mock.Anything).Return(true, nil).Once()

		service := _orderService.NewService(mockOrderRepo, mockCarepo, mockUserRepo)
		order, err := service.Done(mockOrderInput)
		assert.NoError(t, err)
		assert.NotNil(t, order)

		mockOrderRepo.AssertExpectations(t)
		mockCarepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockOrderRepo.Mock.On("FindByID", mock.Anything).Return(mockOrder, nil).Once()
		mockOrderRepo.Mock.On("Done", mock.Anything).Return(false, errors.New("unexpected")).Once()

		service := _orderService.NewService(mockOrderRepo, mockCarepo, mockUserRepo)

		_, err := service.Done(mockOrderInput)

		assert.Error(t, err)

		mockOrderRepo.AssertExpectations(t)
		mockCarepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestOrderService_FindAll(t *testing.T) {
	mockOrderRepo := new(mocks.OrderRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockCarepo := new(mocks.CarRepository)

	mockArrorder := &helpers.Pagination{
		Page:       1,
		Limit:      1,
		Sort:       "id asc",
		TotalRows:  2,
		TotalPages: 2,
		Rows: domain.Order{
			ID:            1,
			EstimatedDays: 3,
			WithDriver:    true,
			Status:        "pending",
			CarID:         1,
			UserID:        1,
			UpdatedAt:     time.Now(),
			CreatedAt:     time.Now(),
		},
	}
	mockEmptyOrder := &helpers.Pagination{}
	mockPagination := helpers.Pagination{
		Page:  1,
		Limit: 1,
		Sort:  "id asc",
	}
	mockEmptyPagination := helpers.Pagination{}

	t.Run("success", func(t *testing.T) {
		mockOrderRepo.Mock.On("FindAll", mock.Anything).Return(mockArrorder, nil).Once()
		service := _orderService.NewService(mockOrderRepo, mockCarepo, mockUserRepo)
		orders, err := service.FindAll(mockPagination)
		assert.NoError(t, err)
		assert.NotNil(t, orders)
		mockOrderRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockOrderRepo.On("FindAll", mock.Anything).Return(mockEmptyOrder, errors.New("Unexpected")).Once()

		service := _orderService.NewService(mockOrderRepo, mockCarepo, mockUserRepo)
		orders, err := service.FindAll(mockEmptyPagination)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyOrder, orders)

		mockOrderRepo.AssertExpectations(t)
	})
}
