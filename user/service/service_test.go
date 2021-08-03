package service_test

import (
	"errors"
	"os"
	auth "section9/auth/service/mocks"
	"section9/domain"
	"section9/domain/mocks"
	"section9/user/input"
	_userService "section9/user/service"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_Create(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockAuthService := new(auth.Service)
	mockUser := &domain.User{
		FullName: "arnold",
		Address:  "california street",
		Mobile:   "08121348584",
		Email:    "arnold@mail.com",
		IDCard:   "id card",
		Role:     "employee",
		Password: "secret",
	}
	mockEmptyUser := &domain.User{}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.Mock.On("FindByEmail", mock.Anything).Return(mockUser, nil).Once()
		mockUserRepo.Mock.On("Create", mock.Anything).Return(mockUser, nil).Once()
		service := _userService.NewService(mockUserRepo, mockAuthService)
		user, err := service.Create(mockUser)
		assert.NoError(t, err)
		assert.NotNil(t, user)

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.Mock.On("FindByEmail", mock.Anything).Return(mockUser, nil).Once()
		mockUserRepo.On("Create", mock.Anything).Return(mockEmptyUser, errors.New("Unexpected")).Once()

		service := _userService.NewService(mockUserRepo, mockAuthService)

		_, err := service.Create(mockEmptyUser)
		assert.Error(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_FindAll(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockAuthService := new(auth.Service)
	mockArruser := []*domain.User{
		&domain.User{
			FullName: "arnold",
			Address:  "california street",
			Mobile:   "08121348584",
			Email:    "arnold@mail.com",
			IDCard:   "id card",
			Role:     "employee",
		},
	}
	mockEmptyUser := []*domain.User{
		&domain.User{},
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.Mock.On("FindAll", mock.Anything).Return(mockArruser, nil).Once()
		service := _userService.NewService(mockUserRepo, mockAuthService)
		users, err := service.FindAll()
		assert.NoError(t, err)
		assert.NotNil(t, users)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("FindAll", mock.Anything).Return(mockEmptyUser, errors.New("Unexpected")).Once()

		service := _userService.NewService(mockUserRepo, mockAuthService)

		users, err := service.FindAll()

		assert.Error(t, err)
		assert.Equal(t, mockEmptyUser, users)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_Update(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockAuthService := new(auth.Service)
	mockUserInput := &input.UpdateInput{
		FullName: "arnold edit",
		Address:  "california street edit",
		Mobile:   "08121348584 edit",
		IDCard:   "id card edit",
	}
	mockUser := &domain.User{
		ID:       1,
		FullName: "arnold",
		Address:  "california street",
		Mobile:   "08121348584",
		Email:    "arnold@mail.com",
		IDCard:   "id card",
		Role:     "employee",
	}
	mockEmptyUser := &input.UpdateInput{}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything).Return(mockUser, nil).Once()
		mockUserRepo.On("Update", mock.Anything, mock.Anything).Return(true, nil).Once()
		service := _userService.NewService(mockUserRepo, mockAuthService)
		isUpdated, err := service.Update(1, mockUserInput)

		assert.NoError(t, err)
		assert.Equal(t, isUpdated, true)
		assert.NotNil(t, isUpdated)

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything).Return(mockUser, nil).Once()
		mockUserRepo.On("Update", mock.Anything).Return(false, errors.New("unexpected")).Once()

		service := _userService.NewService(mockUserRepo, mockAuthService)
		isUpdated, err := service.Update(1, mockEmptyUser)
		assert.Error(t, err)
		assert.Equal(t, isUpdated, false)

		mockUserRepo.AssertExpectations(t)
	})

}

func TestUserService_FindByID(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockAuthService := new(auth.Service)
	mockUser := &domain.User{
		ID:        1,
		FullName:  "arnold",
		Address:   "california street",
		Mobile:    "08121348584",
		Email:     "arnold@mail.com",
		IDCard:    "id card",
		Role:      "employee",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	mockEmptyUser := &domain.User{}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything).Return(mockUser, nil).Once()
		service := _userService.NewService(mockUserRepo, mockAuthService)
		user, err := service.FindByID(1)

		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)
		assert.NotNil(t, user)

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything).Return(mockEmptyUser, errors.New("unexpected")).Once()

		service := _userService.NewService(mockUserRepo, mockAuthService)
		user, err := service.FindByID(1)
		assert.Error(t, err)
		assert.Equal(t, user, mockEmptyUser)

		mockUserRepo.AssertExpectations(t)
	})

}

func TestUserService_Delete(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockAuthService := new(auth.Service)
	mockUser := &domain.User{
		ID:        1,
		FullName:  "arnold",
		Address:   "california street",
		Mobile:    "08121348584",
		Email:     "arnold@mail.com",
		IDCard:    "id card",
		Role:      "employee",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything).Return(mockUser, nil).Once()
		mockUserRepo.On("Delete", mock.Anything).Return(true, nil).Once()
		service := _userService.NewService(mockUserRepo, mockAuthService)
		isDeleted, err := service.Delete(1)

		assert.NoError(t, err)
		assert.Equal(t, isDeleted, true)
		assert.NotNil(t, isDeleted)

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything).Return(mockUser, nil).Once()
		mockUserRepo.On("Delete", mock.Anything).Return(false, errors.New("unexpected")).Once()
		service := _userService.NewService(mockUserRepo, mockAuthService)
		isDeleted, err := service.Delete(1)
		assert.Error(t, err)
		assert.Equal(t, isDeleted, false)

		mockUserRepo.AssertExpectations(t)
	})

}

func TestUserService_Login(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockAuthService := new(auth.Service)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.MinCost)
	mockUser := &domain.User{
		ID:       1,
		FullName: "arnold",
		Address:  "california street",
		Mobile:   "08121348584",
		Email:    "arnold@mail.com",
		IDCard:   "id card",
		Role:     "employee",
		Password: string(passwordHash),
	}
	mockUserInput := input.LoginInput{
		Email:    "arnold@mail.com",
		Password: "secret",
	}

	mockEmptyUser := &domain.User{}

	t.Run("success", func(t *testing.T) {
		access_token := generateToken(mockUser)
		mockUserRepo.Mock.On("FindByEmail", mock.Anything).Return(mockUser, nil).Once()
		mockAuthService.Mock.On("GenerateToken", mock.Anything).Return(access_token, nil).Once()

		service := _userService.NewService(mockUserRepo, mockAuthService)
		user, err := service.Login(mockUserInput)
		assert.NoError(t, err)
		assert.NotNil(t, user)

		mockUserRepo.AssertExpectations(t)
		mockAuthService.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		access_token := generateToken(mockUser)
		mockUserRepo.Mock.On("FindByEmail", mock.Anything).Return(mockUser, nil).Once()
		mockAuthService.Mock.On("GenerateToken", mock.Anything).Return(access_token, nil).Once()

		service := _userService.NewService(mockUserRepo, mockAuthService)

		_, err := service.Create(mockEmptyUser)
		assert.Error(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}

func generateToken(user *domain.User) string {
	claim := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	var SECRET_KEY = os.Getenv("SECRET")

	access_token, _ := token.SignedString([]byte(SECRET_KEY))
	return access_token
}
