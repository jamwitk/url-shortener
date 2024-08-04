package controllers

import (
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func CreateShortUrl(c *gin.Context) {
	id := c.Param("id")
	var shortUrl models.ShortUrl

	fmt.Println("id: ", id)

	err := c.BindJSON(&shortUrl)
	if err != nil {
		c.JSON(
			400,
			gin.H{"error": err.Error()},
		)
		return
	}
	shortUrl.ShortCode = id
	// get application domain

	shortUrl.ShortUrl = c.Request.Host + ShortGenerator()
	c.JSON(200, shortUrl)
}

func ShortGenerator() string {
	Alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var shortUrl = "/"
	for i := 0; i < 6; i++ {
		shortUrl += string(Alphabet[rand.Intn(len(Alphabet))])
	}
	return shortUrl
}
