package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// PathLogger - Log out the path of the request
func PathLogger(c *gin.Context) {
	path := c.FullPath()
	fmt.Println(path)
}
