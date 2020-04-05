package usersController

import (
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/tomvu/poke/db"
	userModel "github.com/tomvu/poke/models/user"
)

func GetAUserHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	user_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	path := c.FullPath()

	userModel.GetAUser(db, user_id)
	message := "requested "

	// c.String(http.StatusOK, "Hello %s", name)
	c.JSON(200, gin.H{
		"message": message,
		"path":    path,
		// "user":    user,
	})
}
