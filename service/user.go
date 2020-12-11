package service

import (
	"go-proj/cmd/app_error"
	"go-proj/domain"
)

type UserService interface {
	FindAll() ([]domain.User, *app_error.AppError)
	FindById(id uint) (*domain.User, *app_error.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) FindAll() ([]domain.User, *app_error.AppError) {
	return s.repo.FindAll()
}

func (s DefaultUserService) FindById(id uint) (*domain.User, *app_error.AppError) {
	return s.repo.FindById(id)
}

func NewUserService(r domain.UserRepository) DefaultUserService {
	return DefaultUserService{r}
}
