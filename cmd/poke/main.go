package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	greet "github.com/tomvu/poke/pkg/greet"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "zxc321"
	dbname   = "poke_development"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	greet.Hello()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		message := "Hello " + name
		path := c.FullPath()
		// c.String(http.StatusOK, "Hello %s", name)
		c.JSON(200, gin.H{
			"message": message,
			"path":    path,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
