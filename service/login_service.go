package service

import "github.com/riwandylbs/go-learning-gin-gonic/repository"

type LoginService interface {
	Login(username string, pwd string) bool
}

type loginService struct {
	userRepository repository.UserRepository
}

func NewLoginService(r repository.UserRepository) *loginService {
	return &loginService{
		userRepository: r,
	}
}

func (l *loginService) Login(username string, pwd string) bool {
	return true
}
