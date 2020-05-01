package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	// TODO: godotenv is currently not used but might be needed in the future
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	pingController "github.com/tomvu/poke/controllers/ping"
	todoscontroller "github.com/tomvu/poke/controllers/todos"
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
	setEnvVars()

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
		users.PUT("/:id", userscontroller.UpdateHandler)
		users.DELETE("/:id", userscontroller.DeleteHandler)
		// TODO: Use PATCH request instead of PUT in the future
		// users.PATCH("/:id", userscontroller.UpdateAUserHandler)

		todos := users.Group("/:id/todos")
		{
			todos.GET("", todoscontroller.GetAllHandler)
			todos.GET("/:todo_id", todoscontroller.GetOneHandler)
			todos.POST("/create", todoscontroller.CreateHandler)
			todos.PUT("/:todo_id", todoscontroller.UpdateHandler)
			todos.DELETE("/:todo_id", todoscontroller.DeleteHandler)
		}
	}

	// Public Group route
	public := router.Group("")
	{
		public.POST("/login", userscontroller.LoginHandler)
		public.POST("/signup", userscontroller.CreateHandler)
	}

	router.GET("/welcome", WelcomeHandler)

	router.GET("/getb", GetDataB)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// router.Run(":3000") // listen and serve on a specified port
}

func setEnvVars() {
	os.Setenv("SECRET_KEY", "Somethingveryimportantmbidk")
}

func WelcomeHandler(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")

	c.JSON(200, gin.H{
		"message": "Welcome " + firstname + " " + lastname,
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
