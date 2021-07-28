package services

import (
	"z-document/models"
	"z-document/repositories"
)

type ArticleService interface {
	CreateArticle(usedId uint) (models.Article, error)
	GetArticleByID(id uint) (models.Article, error)
	GetArticlesByUserId(usedId uint) ([]models.Article, error)
	UpdateArticle(article models.Article) error
	DeleteArticleById(id uint) error
}

type articleService struct {
	repository repositories.ArticleRepository
}

func NewArticleService(articleRepository repositories.ArticleRepository) ArticleService {
	return &articleService{
		repository: articleRepository,
	}
}

func (s *articleService) CreateArticle(usedId uint) (models.Article, error) {
	panic("implement me")
}

func (s *articleService) GetArticleByID(id uint) (models.Article, error) {
	panic("implement me")
}

func (s *articleService) GetArticlesByUserId(usedId uint) ([]models.Article, error) {
	panic("implement me")
}

func (s *articleService) UpdateArticle(article models.Article) error {
	panic("implement me")
}

func (s *articleService) DeleteArticleById(id uint) error {
	panic("implement me")
}
