package service

import (
	"github.com/SaviolaX/blog/internal/dto"
	"github.com/SaviolaX/blog/internal/model"
	"github.com/SaviolaX/blog/internal/repository"
)

type CategoryService interface {
	Create(req *dto.CreateCategoryRequest) error
	FindAll() ([]*dto.CategoryResponse, error)
	FindByID(id uint) (*dto.CategoryResponse, error)
	Delete(id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func (cs *categoryService) FindByID(id uint) (*dto.CategoryResponse, error) {
	category, err := cs.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	var categoryResp dto.CategoryResponse

	categoryResp.ID = category.ID
	categoryResp.Name = category.Name

	return &categoryResp, nil
}

func (cs *categoryService) FindAll() ([]*dto.CategoryResponse, error) {
	categories, err := cs.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesResp []*dto.CategoryResponse
	for _, category := range categories {
		categoriesResp = append(categoriesResp, &dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}
	return categoriesResp, nil
}

func (cs *categoryService) Delete(id uint) error {
	err := cs.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (cs *categoryService) Create(req *dto.CreateCategoryRequest) error {
	var category model.Category

	category.Name = req.Name

	_, err := cs.repo.Create(&category)
	if err != nil {
		return err
	}
	return nil
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}
