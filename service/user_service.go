package service

import (
	"github.com/riwandylbs/go-learning-gin-gonic/models"
	"github.com/riwandylbs/go-learning-gin-gonic/repository"
)

type UserService interface {
	GetAll() ([]models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) *userService {
	return &userService{
		userRepository: r,
	}
}

func (u *userService) GetAll() ([]models.User, error) {
	return u.userRepository.GetAll()
}
