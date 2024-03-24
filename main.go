package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/docs", func(c *gin.Context) {
		c.File("docs/output_file.html")
	})

	r.GET("/api/v1/resource", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.Run(":8080")
}
