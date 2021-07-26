package routers

import (
	"github.com/gin-gonic/gin"
	"z-document/controllers"
	_ "z-document/docs"
	"z-document/repositories"
	"z-document/services"
)

func InitAccountRouter(r *gin.RouterGroup) {
	repository := repositories.NewAccountRepository()
	accountService := services.NewAccountService(repository)
	c := controllers.NewAccountController(accountService)
	auth := r.Group("/auth")
	{
		auth.POST("/login", c.Login)
		auth.POST("/register", c.Register)
	}
}
