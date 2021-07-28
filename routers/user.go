package routers

import (
	"github.com/gin-gonic/gin"
	"z-document/controllers"
	_ "z-document/docs"
	"z-document/repositories"
	"z-document/services"
)

func InitUserRouter(r *gin.RouterGroup) {
	repository := repositories.NewUserRepository()
	userService := services.NewUserService(repository)
	c := controllers.NewUserController(userService)
	auth := r.Group("/auth")
	{
		auth.POST("/login", c.Login)
		auth.POST("/register", c.Register)
	}
}
