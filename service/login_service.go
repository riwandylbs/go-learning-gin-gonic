package service

import (
	"log"

	"github.com/riwandylbs/go-learning-gin-gonic/models"
	"github.com/riwandylbs/go-learning-gin-gonic/repository"
)

type LoginService interface {
	Login(username string, pwd string) ([]models.User, error)
}

type loginService struct {
	userRepository repository.UserRepository
}

func NewLoginService(r repository.UserRepository) *loginService {
	return &loginService{
		userRepository: r,
	}
}

func (l *loginService) Login(username string, pwd string) ([]models.User, error) {
	users, err := l.userRepository.Login(username, pwd)
	if err != nil {
		log.Fatal("terjadi kesalahan ", err)
	}
	return users, err
}
