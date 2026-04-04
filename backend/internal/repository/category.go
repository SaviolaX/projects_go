package repository

import (
	"errors"
	"log"

	"github.com/SaviolaX/blog/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *model.Category) (*model.Category, error)
	FindByID(id uint) (*model.Category, error)
	FindAll() ([]*model.Category, error)
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func (cr *categoryRepository) FindAll() ([]*model.Category, error) {
	var categories []*model.Category

	err := cr.db.Find(&categories)
	if err.Error != nil {
		return categories, errors.New("categories not found")
	}
	return categories, nil
}

func (cr *categoryRepository) Delete(id uint) error {
	err := cr.db.Delete(&model.Category{}, id)
	if err.Error != nil {
		return errors.New("category not deleted")
	}
	return nil
}

func (cr *categoryRepository) FindByID(id uint) (*model.Category, error) {
	var category model.Category

	err := cr.db.First(&category, id)
	if err.Error != nil {
		return nil, errors.New("category not found by ID")
	}

	return &category, nil
}

func (cr *categoryRepository) Create(category *model.Category) (*model.Category, error) {
	log.Println("start to create a new category")
	err := cr.db.Create(&category)
	if err.Error != nil {
		return nil, errors.New("category not created")
	}
	log.Println("category created:", category)
	return category, nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
