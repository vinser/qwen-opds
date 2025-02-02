package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer(port int) {
	r := gin.Default()
	r.GET("/opds", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Qwen OPDS Server",
		})
	})

	log.Printf("Starting server on port %d...\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
