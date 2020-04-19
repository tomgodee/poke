package usersController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/tomvu/poke/db"
	userModel "github.com/tomvu/poke/models/user"
)

func GetOne(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	path := c.FullPath()

	user := userModel.GetOne(db, userID)
	message := "requested "

	c.JSON(200, gin.H{
		"message": message,
		"path":    path,
		"user":    user,
	})
}

func GetAll(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	users := userModel.GetAll(db)
	c.JSON(200, users)
}

func Create(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	newUserData := c.PostFormMap("payload")

	newUserID := userModel.Create(db, newUserData)
	c.JSON(200, newUserID)
}

func Update(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	updatedData := c.PostFormMap("payload")

	userModel.Update(db, updatedData, userID)

	c.JSON(http.StatusOK, "Success")
}

func Delete(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	userModel.Delete(db, userID)

	c.JSON(http.StatusOK, "Success")
}
