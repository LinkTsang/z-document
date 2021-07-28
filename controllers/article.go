package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"z-document/services"
)

type ArticleController interface {
	AddArticle(*gin.Context)
	GetArticle(*gin.Context)
	ListArticles(*gin.Context)
	UpdateArticle(*gin.Context)
	DeleteArticle(*gin.Context)
}

type articleController struct {
	service services.ArticleService
}

func NewArticleController(service services.ArticleService) ArticleController {
	return &articleController{
		service: service,
	}
}

func (c *articleController) AddArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, RespNotImplemented)
}

func (c *articleController) GetArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, RespNotImplemented)
}

func (c *articleController) ListArticles(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, RespNotImplemented)
}

func (c *articleController) UpdateArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, RespNotImplemented)
}

func (c *articleController) DeleteArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, RespNotImplemented)
}
