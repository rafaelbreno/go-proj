package service

import (
	"go-proj/domain"
)

type UserService interface {
	FindAll() ([]domain.User, error)
	FindById(id uint) (domain.User, error)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) FindAll() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s DefaultUserService) FindById(id uint) (domain.User, error) {
	return s.repo.FindById(id)
}

func NewUserService(r domain.UserRepository) DefaultUserService {
	return DefaultUserService{r}
}
