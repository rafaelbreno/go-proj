package service

import (
	"go-proj/domain"
)

type UserService interface {
	GetAllUsers() ([]domain.User, error)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}

func NewUserService(r domain.UserRepository) DefaultUserService {
	return DefaultUserService{r}
}
