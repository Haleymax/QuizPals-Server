package controllers

import (
	"QuizPals-Server/internal/app/services"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type OpenAIController struct {
	openaiServe services.OpenAIService
}

func NewUserController(OpenAIService services.OpenAIService) *OpenAIController {
	return &OpenAIController{
		openaiServe: OpenAIService,
	}
}

func (oc *OpenAIController) UploadFile(c *gin.Context) {
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

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".md" {
		log.Println("File type not supported", ext)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Only supports Markdown files(.md, .markdown)",
			"data": nil,
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		log.Println("Open file error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Unable to read file contents : " + err.Error(),
			"data": nil,
		})
		return
	}
	defer src.Close()

	content, err := io.ReadAll(src)
	if err != nil {
		log.Println("Read file error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Unable to read file contents : " + err.Error(),
			"data": nil,
		})
		return
	}

	question, err := oc.openaiServe.GenerateQuestions(string(content))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Generate question failed : " + err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": gin.H{
			"filename": file.Filename,
			"size":     file.Size,
			"content":  question,
		},
	})
}
