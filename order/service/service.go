package service

import (
	"errors"
	"fmt"
	"section9/domain"
	"section9/order/formatter"
	"section9/order/helpers"
	"section9/order/input"
	"time"
)

type service struct {
	repository     domain.OrderRepository
	carRepository  domain.CarRepository
	userRepository domain.UserRepository
}

func NewService(repository domain.OrderRepository, carRepository domain.CarRepository, userRepository domain.UserRepository) *service {
	return &service{repository, carRepository, userRepository}
}

func (s *service) Create(order *input.Order) (*formatter.Formatter, error) {
	user, err := s.userRepository.FindByID(order.UserID)
	if err != nil {
		return &formatter.Formatter{}, errors.New("internal server error")
	}
	if user.ID == 0 {
		return &formatter.Formatter{}, errors.New("user not found")
	}
	car, _ := s.carRepository.FindByID(order.CarID)
	if car.ID == 0 {
		return &formatter.Formatter{}, errors.New("car not found")
	}
	if !car.IsActive {
		return &formatter.Formatter{}, errors.New("car is not active")
	}

	if car.Stock <= 0 {
		return &formatter.Formatter{}, errors.New("car stock is empty")
	}

	newOrder := domain.Order{
		EstimatedDays: order.EstimatedDays,
		WithDriver:    order.WithDriver,
		Status:        "pending",
		CarID:         order.CarID,
		UserID:        order.UserID,
	}

	driverPrice := 150000
	withDriverPrice := 0
	if order.WithDriver {
		withDriverPrice = driverPrice * order.EstimatedDays
	}
	estimatedPrice := (car.PricePerDay * order.EstimatedDays) + withDriverPrice
	createdOrder, err := s.repository.Create(&newOrder)
	if err != nil {
		return &formatter.Formatter{}, errors.New("failed to insert order to DB")
	}

	// ! minus car stock
	car.Stock -= 1
	_, err = s.carRepository.Update(car)
	if err != nil {
		return &formatter.Formatter{}, errors.New("failed to minus car stock")
	}

	orderID := fmt.Sprintf("order-%d", createdOrder.ID)
	orderFormat := formatter.Formatter{
		OrderID:        orderID,
		Fullname:       user.FullName,
		IDCard:         user.IDCard,
		Car:            car.Label,
		CarType:        car.CarType,
		EstimatedPrice: estimatedPrice,
	}
	return &orderFormat, nil
}

func (s *service) Done(input *input.DoneInput) (*formatter.DoneFormatter, error) {
	order, err := s.repository.FindByID(input.OrderID)
	if err != nil {
		return &formatter.DoneFormatter{}, errors.New("internal server error")
	}

	if order.ID == 0 {
		return &formatter.DoneFormatter{}, errors.New("order not found")
	}

	driverPrice := 150000
	withDriverPrice := 0

	// ! get current year,month,day
	t := time.Now()
	yearNow := t.Year()        // type int
	monthNow := int(t.Month()) // type time.Month
	dayNow := t.Day()

	t1 := Date(yearNow, monthNow, dayNow)

	// ! get created at year,month, day
	orderTime := order.CreatedAt
	orderYear := orderTime.Year()        // type int
	orderMonth := int(orderTime.Month()) // type time.Month
	orderDay := orderTime.Day()
	t2 := Date(orderYear, orderMonth, orderDay)
	days := int(t2.Sub(t1).Hours() / 24)
	if days == 0 {
		days = 1
	}

	if order.WithDriver {
		withDriverPrice = driverPrice * days
	}
	totalPrice := (order.Car.PricePerDay * days) + withDriverPrice

	order.Status = "done"

	_, err = s.repository.Done(order)
	if err != nil {
		return &formatter.DoneFormatter{}, errors.New("failed to update order status")
	}

	orderID := fmt.Sprintf("order-%d", order.ID)
	orderFormat := formatter.DoneFormatter{
		OrderID:    orderID,
		Fullname:   order.User.FullName,
		IDCard:     order.User.IDCard,
		Car:        order.Car.Label,
		CarType:    order.Car.CarType,
		TotalPrice: totalPrice,
	}

	return &orderFormat, nil
}
func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (s *service) FindAll(pagination helpers.Pagination) (*helpers.Pagination, error) {
	datas, err := s.repository.FindAll(pagination)
	if err != nil {
		return datas, err
	}
	return datas, nil
}
