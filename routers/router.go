package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"z-document/controllers"
	_ "z-document/docs"
)

func InitRouter() *gin.Engine {
	port := 8000
	r := gin.Default()
	c := controllers.NewController()

	r.GET("/", c.Index)

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port)) // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api := r.Group("/api")

	InitAccountRouter(api)

	return r
}
