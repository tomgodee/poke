package pingController

import "github.com/gin-gonic/gin"

// Ping API handler
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}
