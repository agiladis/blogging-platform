package repository

import "blogging-platform/model"

type BlogPostRepository interface {
	Create(blogPost model.BlogPost) (model.BlogPost, error)
	GetAll() ([]model.BlogPost, error)
	GetById(blogPostId int) (model.BlogPost, error)
	Update(blogPostId int, blogPost model.BlogPost) error
	Delete(blogPostId int) error
}
