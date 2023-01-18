package repository

import (
	"log"

	"github.com/riwandylbs/go-learning-gin-gonic/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	Save(models.User) (models.User, error)
	Login(username string, password string) ([]models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) GetAll() (users []models.User, err error) {
	log.Print("Get all users from database")
	err = u.DB.Find(&users).Error
	if err != nil {
		log.Fatalf("error when get data from database")
	}
	return users, err
}

func (u *userRepository) Save(user models.User) (models.User, error) {
	return user, nil
}

func (u *userRepository) Login(username string, pwd string) ([]models.User, error) {
	var users []models.User

	result := u.DB.Where("email = ? AND password = ?", "a@email.com", "123456").Find(&users).Error
	if result != nil {
		log.Fatal("errors : ", result)
		return nil, result
	}

	return users, nil
}
