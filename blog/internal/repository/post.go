package repository

import (
	"errors"

	"github.com/SaviolaX/blog/internal/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
	FindAll(limit, offset int) (int64, []model.Post, error)
	FindByID(id uint) (*model.Post, error)
	Update(post *model.Post) error
	Delete(id uint) error
	FindByCategoryID(id uint) ([]model.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func (pr *postRepository) FindByCategoryID(id uint) ([]model.Post, error) {
	var posts []model.Post

	err := pr.db.Preload("Category").Preload("Author").Order("created_at DESC").Where("category_id = ?", id).Find(&posts)
	if err.Error != nil {
		return posts, errors.New("posts not found")
	}
	return posts, nil
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

	err := pr.db.Preload("Category").Preload("Author").First(&post, id)
	if err.Error != nil {
		return nil, errors.New("post not found by ID")
	}

	return &post, nil
}

func (pr *postRepository) FindAll(limit, offset int) (int64, []model.Post, error) {
	var posts []model.Post
	var totalPosts int64

	pr.db.Model(&model.Post{}).Count(&totalPosts)

	err := pr.db.Preload("Category").Preload("Author").Order("created_at DESC").Limit(limit).Offset(offset).Find(&posts)
	if err.Error != nil {
		return 0, posts, errors.New("posts not found")
	}

	return totalPosts, posts, nil
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
