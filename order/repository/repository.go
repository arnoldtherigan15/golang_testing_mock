package repository

import (
	"math"
	"section9/domain"
	"section9/order/helpers"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	db.AutoMigrate(&domain.Order{})
	return &repository{db}
}

func (r *repository) Create(order *domain.Order) (*domain.Order, error) {
	err := r.db.Create(order).Error
	if err != nil {
		return &domain.Order{}, err
	}
	return order, nil
}

func (r *repository) Done(order *domain.Order) (bool, error) {
	err := r.db.Save(order).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) FindByID(ID int) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Car").Preload("Customer").Where("id = ?", ID).Find(&order).Error
	if err != nil {
		return &order, err
	}
	return &order, nil
}

func (r *repository) FindAll(pagination helpers.Pagination) (*helpers.Pagination, error) {
	var orders []*domain.Order
	err := r.db.Scopes(paginate(orders, &pagination, r.db)).Preload("Car").Preload("Customer").Find(&orders).Error
	pagination.Rows = orders
	// err := r.db.Find(&orders).Error
	if err != nil {
		return &pagination, err
	}
	return &pagination, nil
	// return orders, nil
}

func paginate(value interface{}, pagination *helpers.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
