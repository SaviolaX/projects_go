package dto

import (
	"errors"
	"strings"
	"time"
)

type UpdatePostRequest struct {
	Title      string
	Entry      string
	CategoryID uint
}

type PostResponse struct {
	ID        uint
	Title     string
	Entry     string
	Author    UserResponse
	Category  CategoryResponse
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreatePostRequest struct {
	Title        string
	Entry        string
	CategoryName string
	CategoryID   uint
}

func (rr *CreatePostRequest) Validate() error {
	rr.Title = strings.TrimSpace(rr.Title)
	if len(rr.Title) < 4 {
		return errors.New("incorrect post title")
	}

	rr.Entry = strings.TrimSpace(rr.Entry)
	if len(rr.Entry) < 5 {
		return errors.New("incorrect post entry")
	}

	return nil
}
