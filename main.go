package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	uuid "github.com/uuid6/uuid6go-proto"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		if c.Query("version") == "7" {
			var generator uuid.UUIDv7Generator
			id := generator.Next()
			c.String(http.StatusOK, id.ToString())
		}
	})

	router.Run(":" + port)
}
