package repositories

import (
	"gorm.io/gorm"
	"z-document/db"
	"z-document/models"
)

type ArticleRepository interface {
	Create(article models.Article) error
	FindById(id uint) (models.Article, error)
	FindByUserId(usedId uint) ([]models.Article, error)
	FindAll() ([]models.Article, error)
	Update(article models.Article) error
	DeleteById(id uint) error
	Delete(article models.Article) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository() ArticleRepository {
	db := db.GetMySqlDB()
	err := db.AutoMigrate(&models.Article{})
	if err != nil {
		return nil
	}
	return &articleRepository{
		db: db,
	}
}

func (r *articleRepository) Create(article models.Article) error {
	panic("implement me")
}

func (r *articleRepository) FindById(id uint) (models.Article, error) {
	panic("implement me")
}

func (r *articleRepository) FindByUserId(id uint) ([]models.Article, error) {
	panic("implement me")
}

func (r *articleRepository) FindAll() ([]models.Article, error) {
	panic("implement me")
}
func (r *articleRepository) Update(article models.Article) error {
	panic("implement me")
}

func (r *articleRepository) Delete(article models.Article) error {
	panic("implement me")
}

func (r *articleRepository) DeleteById(id uint) error {
	panic("implement me")
}
