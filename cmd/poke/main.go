package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	pingController "github.com/tomvu/poke/controllers/ping"
	userscontroller "github.com/tomvu/poke/controllers/users"
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

	// Private Group user route
	users := router.Group("/users", middlewares.Authentication)
	{
		users.GET("/:id", userscontroller.GetOneHandler)
		users.GET("", userscontroller.GetAllHandler)
		users.POST("", userscontroller.CreateHandler)
		users.PUT("/:id", userscontroller.UpdateHandler)
		users.DELETE("/:id", userscontroller.DeleteHandler)
		// TODO: Use PATCH request instead of PUT in the future
		// users.PATCH("/:id", userscontroller.UpdateAUserHandler)
	}

	// Public Group user route
	publicUsers := router.Group("/users")
	{
		publicUsers.POST("/login", userscontroller.LoginHandler)
	}

	router.GET("/welcome", WelcomeHandler)

	router.POST("form", BasicFormHandler)

	router.GET("/getb", GetDataB)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// router.Run(":3000") // listen and serve on a specified port
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
