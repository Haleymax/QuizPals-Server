package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type OpenAIController struct {
}

func NewUserController() *OpenAIController {
	return &OpenAIController{}
}

func (ctrl *OpenAIController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("mdFile")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	log.Println(file.Filename)

	//err = c.SaveUploadedFile(file, file.Filename)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": gin.H{
			"filename": file.Filename,
		},
	})
}
