package repository

import (
	"section9/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	db.AutoMigrate(&domain.User{})
	return &repository{db}
}

func (r *repository) Create(user *domain.User) (*domain.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func (r *repository) FindAll() ([]*domain.User, error) {
	var users []*domain.User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) Update(user *domain.User) (bool, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) FindByID(ID int) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (r *repository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (r *repository) Delete(user *domain.User) (bool, error) {
	if err := r.db.Delete(user).Error; err != nil {
		return false, err
	}
	return true, nil
}
