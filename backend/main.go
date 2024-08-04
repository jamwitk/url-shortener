package main

import (
	"backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/shortener/:id", controllers.CreateShortUrl)
	err := r.Run()
	if err != nil {
		return
	}
}
