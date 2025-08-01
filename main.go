package main

import (
	"QuizPals-Server/internal/routers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := gin.Default()

	go func() {
		if err := r.Run(":8080"); err != nil {
			panic("Failed to start server: " + err.Error())
		}
	}()

	routers.SetupRouters(r)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
