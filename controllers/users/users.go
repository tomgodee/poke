package usersController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/tomvu/poke/db"
	userModel "github.com/tomvu/poke/models/user"
)

func GetAUserHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	path := c.FullPath()

	user := userModel.GetAUser(db, userID)
	message := "requested "

	c.JSON(200, gin.H{
		"message": message,
		"path":    path,
		"user":    user,
	})
}

func GetUsersHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	users := userModel.GetUsers(db)
	c.JSON(200, users)
}

func CreateAUserHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	newUserData := c.PostFormMap("payload")

	newUserID := userModel.CreateAUser(db, newUserData)
	c.JSON(200, newUserID)
}

func UpdateAUserHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	updatedData := c.PostFormMap("payload")

	userModel.UpdateAUser(db, updatedData, userID)

	c.JSON(http.StatusOK, "a")
}

func DeleteAUserHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

}
