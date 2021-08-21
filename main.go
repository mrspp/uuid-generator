package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/uuid6/uuid6go-proto"
)

func main() {

	router := gin.Default()

	router.GET("/generate",
		func(c *gin.Context) {
			// if c.Query("version") == "6" {

			// 	id := uuid.UUIDv6FromBytes()
			// 	c.String(http.StatusOK, id.ToString())
			// }

			if c.Query("version") == "7" {
				var generator uuid.UUIDv7Generator
				id := generator.Next()
				c.String(http.StatusOK, id.ToString())
			}

		})

	router.GET("/validate", func(c *gin.Context) {
		fmt.Println("validate")

		c.String(http.StatusOK, c.Param("uuid"))

	})

	router.Run(":8080")
}
