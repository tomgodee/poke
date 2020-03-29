package main

import (
	"fmt"
	"io"
	"os"

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

// Ping API handler
func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}

// Middleware that prints out the path
func pathLogger(c *gin.Context) {
	path := c.FullPath()
	fmt.Println(path)
}

func writeLogFile() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

}

func main() {
	writeLogFile()

	// Default With the Logger and Recovery middleware already attached
	router := gin.Default()

	// Ping API
	router.GET("/ping", pathLogger, pong)

	// Test if call a function from an imported pkg works
	greet.Hello()

	//Group user route
	users := router.Group("/users", pathLogger)
	{
		users.GET("/:name", func(c *gin.Context) {
			name := c.Param("name")
			message := "Hello " + name
			path := c.FullPath()
			// c.String(http.StatusOK, "Hello %s", name)
			c.JSON(200, gin.H{
				"message": message,
				"path":    path,
			})
		})

		users.GET("/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			path := c.FullPath()
			c.JSON(200, gin.H{
				"message": message,
				"path":    path,
			})
		})
	}

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.JSON(200, gin.H{
			"message": "Welcome " + firstname + " " + lastname,
		})
	})

	router.POST("form", func(c *gin.Context) {
		// message := c.PostForm("message")
		message := c.PostFormMap("message")
		nick := c.DefaultPostForm("nick", "guest")

		c.JSON(200, gin.H{
			"message": message,
			"nick":    nick,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// router.Run(":3000") // listen and serve on a specified port
}
