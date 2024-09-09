package service

import (
	"blogging-platform/dto"
	"blogging-platform/model"
	"blogging-platform/repository"
)

type BlogPostService interface {
	Create(dto.BlogPostDTO, uint) (model.BlogPost, error)
}

type blogPostService struct {
	blogPostRepository repository.BlogPostRepository
}

func NewBlogPostService(repo repository.BlogPostRepository) *blogPostService {
	return &blogPostService{blogPostRepository: repo}
}

func (s *blogPostService) Create(blogPostDTO dto.BlogPostDTO,userId uint) (model.BlogPost, error) {
	blogPost := model.BlogPost{
		Title: blogPostDTO.Title,
		Content: blogPostDTO.Content,
		UserID: userId,
	}

	return s.blogPostRepository.Create(blogPost)
}