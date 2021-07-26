package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"z-document/services"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccountController interface {
	Login(*gin.Context)
	Register(*gin.Context)
}

type accountController struct {
	service services.AccountService
	AccountController
}

func NewAccountController(service services.AccountService) AccountController {
	return &accountController{
		service: service,
	}
}

// Register godoc
// @Summary register
// @Description register
// @Accept json
// @Produce json
// @Router /auth/register [post]
func (c *accountController) Register(ctx *gin.Context) {
	var json RegisterRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, RespNotImplemented)
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess)
}

// Login godoc
// @Summary login
// @Description login
// @Accept json
// @Produce json
// @Router /auth/login [post]
func (c *accountController) Login(ctx *gin.Context) {
	var json LoginRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, RespNotImplemented)
		return
	}
	ctx.JSON(http.StatusOK, RespSuccess)
}
