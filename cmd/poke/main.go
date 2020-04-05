package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	pingController "github.com/tomvu/poke/controllers/ping"
	database "github.com/tomvu/poke/db"
	"github.com/tomvu/poke/middlewares"
	"github.com/tomvu/poke/pkg/greet"
)

func writeLogFile() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	// Write actions to a log file as well as the console
	writeLogFile()

	// Default With the Logger and Recovery middleware already attached
	router := gin.Default()

	// Ping API
	router.GET("/ping", middlewares.PathLogger, pingController.Pong)

	// Test if call a function from an imported pkg works
	greet.Hello()

	// Connect to the db
	db := database.Connect()

	defer db.Close()

	//Group user route
	users := router.Group("/users", middlewares.PathLogger)
	{
		users.GET("/:name")
	}

	router.GET("/welcome", WelcomeHandler)

	router.POST("form", BasicFormHandler)

	router.GET("/getb", GetDataB)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// router.Run(":3000") // listen and serve on a specified port
}

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

func WelcomeHandler(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")

	c.JSON(200, gin.H{
		"message": "Welcome " + firstname + " " + lastname,
	})
}

func BasicFormHandler(c *gin.Context) {
	// message := c.PostForm("message")
	message := c.PostFormMap("message")
	nick := c.DefaultPostForm("nick", "guest")

	c.JSON(200, gin.H{
		"message": message,
		"nick":    nick,
	})
}

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}
