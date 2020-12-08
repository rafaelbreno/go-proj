package service

import (
	"go-proj/domain"
)

type UserService interface {
	FindAll() ([]domain.User, error)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) FindAll() ([]domain.User, error) {
	return s.repo.FindAll()
}

func NewUserService(r domain.UserRepository) DefaultUserService {
	return DefaultUserService{r}
}
