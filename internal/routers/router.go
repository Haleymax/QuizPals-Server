package routers

import "github.com/gin-gonic/gin"

func SetupRouters(router *gin.Engine) {
	api := router.Group("/api/v1")
	index := api.Group("/index")
	index.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to QuizPals API",
		})
	})
}
