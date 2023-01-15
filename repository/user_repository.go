package repository

import (
	"log"

	"github.com/riwandylbs/go-learning-gin-gonic/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	Save(models.User) (models.User, error)
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
