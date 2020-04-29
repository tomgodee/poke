package todoscontroller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/tomvu/poke/db"
	todomodel "github.com/tomvu/poke/models/todo"
)

func GetAllHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	userIDstr, _ := c.Params.Get("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		panic(err)
	}
	todos := todomodel.GetAll(db, userID)

	c.JSON(200, todos)
}

func GetOneHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	todoIDstr, _ := c.Params.Get("todo_id")
	todoID, err := strconv.Atoi(todoIDstr)
	if err != nil {
		panic(err)
	}
	todo := todomodel.GetOne(db, todoID)

	c.JSON(200, todo)
}

func CreateHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	body := c.PostForm("body")
	userIDstr, _ := c.Params.Get("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		panic(err)
	}
	todomodel.Create(db, body, userID)
	c.JSON(200, "Success")
}

func UpdateHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	payload := c.PostFormMap("payload")
	todoIDstr, _ := c.Params.Get("todo_id")
	todoID, err := strconv.Atoi(todoIDstr)

	if err != nil {
		panic(err)
	}
	todomodel.Update(db, payload, todoID)
	c.JSON(200, "Success")
}

func DeleteHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	todoIDstr, _ := c.Params.Get("todo_id")
	todoID, err := strconv.Atoi(todoIDstr)
	if err != nil {
		panic(err)
	}
	todomodel.Delete(db, todoID)

	c.JSON(200, "Success")
}
