package usersController

import "github.com/gin-gonic/gin"

func GetAUserHandler(c *gin.Context) {
	name := c.Param("name")
	message := "Hello " + name
	path := c.FullPath()
	// c.String(http.StatusOK, "Hello %s", name)
	c.JSON(200, gin.H{
		"message": message,
		"path":    path,
	})
}
