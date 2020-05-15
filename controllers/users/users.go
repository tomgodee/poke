package userscontroller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/tomvu/poke/db"
	usermodel "github.com/tomvu/poke/models/user"
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

	user := usermodel.GetOne(db, userID)
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

	users := usermodel.GetAll(db)
	c.JSON(200, users)
}

func CreateHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	newUserData := c.PostFormMap("payload")

	newUserID := usermodel.Create(db, newUserData)
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

	usermodel.Update(db, updatedData, userID)

	c.JSON(http.StatusOK, "Success")
}

func DeleteHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	usermodel.Delete(db, userID)

	c.JSON(http.StatusOK, "Success")
}

func LoginHandler(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	var user usermodel.User
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	token, err := usermodel.Login(db, user)

	switch {
	case err == sql.ErrNoRows:
		fmt.Print("username")
		c.JSON(401, "Wrong username!")
	case err == bcrypt.ErrMismatchedHashAndPassword:
		fmt.Print("password")
		c.JSON(401, "Wrong password!")
	case err == nil:
		// TODO: What if i wanna set or send token here ?
		// Also need to send back the user the client here
		// c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		c.JSON(200, token)
	}
}
