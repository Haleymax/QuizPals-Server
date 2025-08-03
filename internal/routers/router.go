package routers

import (
	"QuizPals-Server/internal/app/controllers"
	"QuizPals-Server/internal/app/services"
	"github.com/gin-gonic/gin"
)

func SetupRouters(router *gin.Engine) {

	openaiService := services.NewOpenAIService()

	openaiController := controllers.NewUserController(openaiService)

	api := router.Group("/api/v1")
	index := api.Group("/index")

	index.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to QuizPals API",
		})
	})

	openai := api.Group("/openai")
	{
		openai.POST("/upload", openaiController.UploadFile)
	}

}
