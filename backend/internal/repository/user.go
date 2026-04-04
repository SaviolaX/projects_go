package repository

import (
	"errors"

	"github.com/SaviolaX/blog/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User

	err := ur.db.First(&user, "username = ?", username)
	if err.Error != nil {
		return nil, errors.New("user not found by username")
	}

	return &user, nil
}

func (ur *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User

	err := ur.db.First(&user, id)
	if err.Error != nil {
		return nil, errors.New("user not found by ID")
	}

	return &user, nil
}

func (ur *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	err := ur.db.First(&user, "email = ?", email)
	if err.Error != nil {
		return nil, errors.New("user not found by email")
	}

	return &user, nil
}

func (ur *userRepository) Create(user *model.User) error {
	err := ur.db.Create(&user)
	if err.Error != nil {
		return errors.New("user not created")
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
