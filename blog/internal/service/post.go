package service

import (
	"github.com/SaviolaX/blog/internal/dto"
	"github.com/SaviolaX/blog/internal/model"
	"github.com/SaviolaX/blog/internal/repository"
)

type PostService interface {
	Create(req *dto.CreatePostRequest, authorID uint) error
	FindAll() ([]*dto.PostResponse, error)
	FindByID(id uint) (*dto.PostResponse, error)
	Update(id uint, req *dto.UpdatePostRequest) error
	Delete(id uint) error
}

type postService struct {
	repo         repository.PostRepository
	categoryRepo repository.CategoryRepository
}

func (ps *postService) Delete(id uint) error {
	err := ps.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (ps *postService) Update(id uint, req *dto.UpdatePostRequest) error {
	var post model.Post

	post.ID = id
	post.Title = req.Title
	post.Entry = req.Entry

	if req.CategoryName != "" {
		newCategory, err := ps.categoryRepo.Create(&model.Category{
			Name: req.CategoryName,
		})
		if err != nil {
			return err
		}
		post.CategoryID = newCategory.ID
	} else if req.CategoryID != 0 {
		post.CategoryID = req.CategoryID
	}

	err := ps.repo.Update(&post)
	if err != nil {
		return err
	}
	return nil
}

func (ps *postService) FindByID(id uint) (*dto.PostResponse, error) {
	post, err := ps.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.PostResponse{
		ID:    post.ID,
		Title: post.Title,
		Entry: post.Entry,
		Category: dto.CategoryResponse{
			ID:   post.CategoryID,
			Name: post.Category.Name,
		},
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

func (ps *postService) FindAll() ([]*dto.PostResponse, error) {
	posts, err := ps.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var postsResp []*dto.PostResponse
	for _, post := range posts {
		entry := post.Entry
		if len(entry) > 200 {
			entry = entry[:200] + "..."
		}
		postsResp = append(postsResp, &dto.PostResponse{
			ID:    post.ID,
			Title: post.Title,
			Entry: entry,
			Category: dto.CategoryResponse{
				ID:   post.CategoryID,
				Name: post.Category.Name,
			},
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return postsResp, nil
}

func (ps *postService) Create(req *dto.CreatePostRequest, authorID uint) error {
	var post model.Post

	post.Title = req.Title
	post.Entry = req.Entry
	post.AuthorID = authorID

	if req.CategoryID != 0 {
		post.CategoryID = req.CategoryID
	}
	if req.CategoryName != "" {
		newCategory, err := ps.categoryRepo.Create(&model.Category{
			Name: req.CategoryName,
		})
		if err != nil {
			return err
		}
		post.CategoryID = newCategory.ID
	}

	err := ps.repo.Create(&post)
	if err != nil {
		return err
	}
	return nil
}

func NewPostService(repo repository.PostRepository, categoryRepo repository.CategoryRepository) PostService {
	return &postService{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}
