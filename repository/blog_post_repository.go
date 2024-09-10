package repository

import (
	"blogging-platform/model"

	"gorm.io/gorm"
)

type BlogPostRepository interface {
	Create(blogPost model.BlogPost) (model.BlogPost, error)
	GetAll() ([]model.BlogPost, error)
	GetById(blogPostId int) (model.BlogPost, error)
	Update(blogPostId int, blogPost model.BlogPost) error
	Delete(blogPostId int) error
}

type blogPostRepository struct {
	DB *gorm.DB
}

func NewBlogPostRepository(DB *gorm.DB) *blogPostRepository {
	return &blogPostRepository{DB}
}

func (r *blogPostRepository) Create(blogPost model.BlogPost) (model.BlogPost, error) {
	err := r.DB.Create(&blogPost).Error
	return blogPost, err
}
