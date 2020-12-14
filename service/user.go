package service

import (
	"go-proj/cmd/app_error"
	"go-proj/domain"
	"go-proj/dto"
)

type UserService interface {
	FindAll() ([]domain.User, *app_error.AppError)
	FindById(id uint) (*dto.UserResponse, *app_error.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) FindAll() ([]domain.User, *app_error.AppError) {
	return s.repo.FindAll()
}

func (s DefaultUserService) FindById(id uint) (*dto.UserResponse, *app_error.AppError) {
	u, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	ud := dto.UserResponse{
		ID:      u.ID,
		Email:   u.Email,
		Status:  u.Status,
		Account: u.Account,
	}

	return &ud, nil
}

func NewUserService(r domain.UserRepository) DefaultUserService {
	return DefaultUserService{r}
}
