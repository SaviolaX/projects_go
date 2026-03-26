package dto

import (
	"errors"
	"net/mail"
	"strings"
	"time"
)

type UserResponse struct {
	ID        uint
	Username  string
	Email     string
	CreatedAt time.Time
}

type LoginRequest struct {
	Username string
	Password string
}

func (lr *LoginRequest) Validate() error {
	lr.Username = strings.TrimSpace(lr.Username)
	if len(lr.Username) <= 3 {
		return errors.New("incorrect username")
	}

	lr.Password = strings.TrimSpace(lr.Password)
	if len(lr.Password) < 6 {
		return errors.New("incorrect password")
	}

	return nil
}

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

func (rr *RegisterRequest) Validate() error {

	rr.Username = strings.TrimSpace(rr.Username)
	if len(rr.Username) <= 3 {
		return errors.New("incorrect username")
	}

	_, err := mail.ParseAddress(rr.Email)
	if err != nil {
		return errors.New("incorrect email")
	}

	rr.Password = strings.TrimSpace(rr.Password)
	if len(rr.Password) < 6 {
		return errors.New("incorrect password")
	}

	return nil
}
