package userscontroller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/tomvu/poke/db"
	userModel "github.com/tomvu/poke/models/user"
	"golang.org/x/crypto/bcrypt"
)

func GetOneHandler(c *gin.Context) {
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

func GetAllHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	users := userModel.GetAll(db)
	c.JSON(200, users)
}

func CreateHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	newUserData := c.PostFormMap("payload")

	newUserID := userModel.Create(db, newUserData)
	c.JSON(200, newUserID)
}

func UpdateHandler(c *gin.Context) {
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

func DeleteHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	userModel.Delete(db, userID)

	c.JSON(http.StatusOK, "Success")
}

func LoginHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	loginData := c.PostFormMap("payload")
	err := userModel.Login(db, loginData)

	switch {
	case err == sql.ErrNoRows:
		c.JSON(401, "Wrong username!")
	case err == bcrypt.ErrMismatchedHashAndPassword:
		c.JSON(401, "Wrong password!")
	case err == nil:
		c.JSON(200, "Success")
	}
}
