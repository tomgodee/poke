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

	todos := todomodel.GetAll(db)

	c.JSON(200, todos)
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
