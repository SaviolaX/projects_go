package repository

import (
	"errors"
	"log"

	"github.com/SaviolaX/blog/internal/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
	FindAll() ([]model.Post, error)
	FindByID(id uint) (*model.Post, error)
	Update(post *model.Post) error
	Delete(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func (pr *postRepository) Delete(id uint) error {
	err := pr.db.Delete(&model.Post{}, id)
	if err.Error != nil {
		return errors.New("post not deleted")
	}
	return nil
}

func (pr *postRepository) Update(post *model.Post) error {
	err := pr.db.Model(&post).Updates(model.Post{Title: post.Title, Entry: post.Entry, CategoryID: post.CategoryID})
	if err.Error != nil {
		return errors.New("post not updated")
	}
	return nil
}

func (pr *postRepository) FindByID(id uint) (*model.Post, error) {
	var post model.Post

	err := pr.db.Preload("Category").First(&post, id)
	log.Println("repo post category:", post.Category)
	if err.Error != nil {
		return nil, errors.New("post not found by ID")
	}

	return &post, nil
}

func (pr *postRepository) FindAll() ([]model.Post, error) {
	var posts []model.Post

	err := pr.db.Preload("Category").Find(&posts)
	if err.Error != nil {
		return posts, errors.New("posts not found")
	}
	return posts, nil
}

func (pr *postRepository) Create(post *model.Post) error {
	err := pr.db.Create(&post)
	if err.Error != nil {
		return errors.New("post not created")
	}
	return nil
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}
