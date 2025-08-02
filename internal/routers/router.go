package routers

import (
	"QuizPals-Server/internal/app/controllers"
	"QuizPals-Server/internal/app/services"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func SetupRouters(router *gin.Engine) {
	api := router.Group("/api/v1")
	index := api.Group("/index")

	openaiController := controllers.NewUserController()
	index.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to QuizPals API",
		})
	})

}
