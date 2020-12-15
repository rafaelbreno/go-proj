package service

import (
	"go-proj/cmd/app_error"
	"go-proj/domain"
	"go-proj/dto"
)

type UserService interface {
	FindAll() ([]*dto.UserResponse, *app_error.AppError)
	FindById(id uint) (*dto.UserResponse, *app_error.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) FindAll() ([]*dto.UserResponse, *app_error.AppError) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var usersDTO []*dto.UserResponse

	for _, value := range users {
		a := value.ToDTO()
		usersDTO = append(usersDTO, &a)
	}

	return usersDTO, nil
}

func (s DefaultUserService) FindById(id uint) (*dto.UserResponse, *app_error.AppError) {
	u, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	// User DTO
	uDTO := u.ToDTO()

	return &uDTO, nil
}

func NewUserService(r domain.UserRepository) DefaultUserService {
	return DefaultUserService{r}
}
