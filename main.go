package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	// TODO: godotenv is currently not used but might be needed in the future
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	pingController "github.com/tomvu/poke/controllers/ping"
	todoscontroller "github.com/tomvu/poke/controllers/todos"
	userscontroller "github.com/tomvu/poke/controllers/users"
	"github.com/tomvu/poke/middlewares"
)

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsHandler(writer gin.ResponseWriter, request *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	setEnvVars()
	// Default With the Logger and Recovery middleware already attached
	router := gin.Default()
	// router.Static("/assets", "./assets")
	// Ping API
	router.GET("/ping", middlewares.PathLogger, pingController.Pong)

	router.GET("/ws", func(c *gin.Context) {
		WsHandler(c.Writer, c.Request)
	})

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

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// router.Run(":3000") // listen and serve on a specified port
}

func setEnvVars() {
	os.Setenv("SECRET_KEY", "Somethingveryimportantmbidk")
}
