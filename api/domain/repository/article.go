package repository

import (
	"github.com/TadaTeruki/portfolio-server-next/api/domain/entity"
)

type IArticleRepository interface {
	PostArticle(article *entity.PostArticle) (string, error)
	GetArticle(articleID string) (*entity.GetArticle, error)
	DeleteArticle(articleID string) error
	UpdateArticle(articleID string, updatedArticle *entity.UpdateArticle) error
	GetArticles(number int) ([]entity.ListArticle, error)
}
