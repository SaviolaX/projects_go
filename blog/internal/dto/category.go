package dto

import (
	"errors"
	"strings"
)

type CreateCategoryRequest struct {
	Name string
}

type CategoryResponse struct {
	ID   uint
	Name string
}

func (cr *CreateCategoryRequest) Validate() error {
	cr.Name = strings.TrimSpace(cr.Name)
	if len(cr.Name) < 2 {
		return errors.New("incocrect category name")
	}

	return nil
}
