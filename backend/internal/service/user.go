package service

import (
	"errors"

	"github.com/SaviolaX/blog/internal/dto"
	"github.com/SaviolaX/blog/internal/model"
	"github.com/SaviolaX/blog/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func (us *userService) Register(req *dto.RegisterRequest) error {
	_, err := us.repo.FindByEmail(req.Email)
	if err == nil {
		return errors.New("user with this email already exists")
	}

	hashedPasswordStr, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("hashing password error")
	}

	var newUser model.User

	newUser.Username = req.Username
	newUser.Email = req.Email
	newUser.Password = string(hashedPasswordStr)

	err = us.repo.Create(&newUser)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) Login(req *dto.LoginRequest) (*dto.UserResponse, error) {
	var resp dto.UserResponse

	user, err := us.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	resp.ID = user.ID
	resp.Username = user.Username
	resp.Email = user.Email
	resp.CreatedAt = user.CreatedAt

	return &resp, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
