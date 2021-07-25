package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    CodeSuccess,
		"message": "Hello ZNote Go Server!",
	})
}
