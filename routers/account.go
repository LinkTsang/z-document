package routers

import (
	"github.com/gin-gonic/gin"
	"znote-server-go/controllers"
	_ "znote-server-go/docs"
	"znote-server-go/repositories"
	"znote-server-go/services"
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
