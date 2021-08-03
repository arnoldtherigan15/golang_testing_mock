package service

import (
	"errors"
	auth "section9/auth/service"
	"section9/domain"
	"section9/user/formatter"
	"section9/user/input"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository  domain.UserRepository
	authService auth.Service
}

func NewService(repository domain.UserRepository, authService auth.Service) *service {
	return &service{repository, authService}
}

func (s *service) Create(us *domain.User) (*domain.User, error) {
	userExists, err := s.repository.FindByEmail(us.Email)
	if err != nil {
		return &domain.User{}, errors.New("internal server error")
	}

	if userExists.ID != 0 {
		return &domain.User{}, errors.New("email already exists")
	}
	us.Role = "customer"
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.MinCost)
	if err != nil {
		return &domain.User{}, err
	}

	us.Password = string(passwordHash)
	user, err := s.repository.Create(us)
	if err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (s *service) FindAll() ([]*domain.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *service) Update(ID int, user *input.UpdateInput) (bool, error) {
	userDB, err := s.repository.FindByID(ID)
	if err != nil {
		return false, err
	}
	if userDB.ID == 0 {
		return false, errors.New("user not found")
	}
	userDB.FullName = user.FullName
	userDB.Address = user.Address
	userDB.Mobile = user.Mobile
	userDB.IDCard = user.IDCard

	updatedUser, err := s.repository.Update(userDB)
	if err != nil {
		return false, err
	}
	return updatedUser, nil
}

func (s *service) FindByID(ID int) (*domain.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return &domain.User{}, err
	}
	if user.ID == 0 {
		return &domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *service) Delete(ID int) (bool, error) {
	user, err := s.FindByID(ID)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return false, errors.New("user not found")
	}
	isDeleted, err := s.repository.Delete(user)
	if err != nil {
		return false, err
	}
	return isDeleted, nil
}

func (s *service) Login(input input.LoginInput) (*formatter.LoginResponse, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return &formatter.LoginResponse{}, err
	}
	if user.ID == 0 {
		return &formatter.LoginResponse{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return &formatter.LoginResponse{}, errors.New("invalid email or password")
	}

	access_token, err := s.authService.GenerateToken(user)

	if err != nil {
		return &formatter.LoginResponse{}, errors.New("internal server error")
	}

	response := formatter.LoginResponse{
		ID:          user.ID,
		Email:       user.Email,
		AccessToken: access_token,
	}
	return &response, nil
}
