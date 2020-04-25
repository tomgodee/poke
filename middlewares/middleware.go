package middlewares

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// PathLogger - Log out the path of the request
func PathLogger(c *gin.Context) {
	path := c.FullPath()
	fmt.Println(path)
}

// Authentication - Validate incoming token
func Authentication(c *gin.Context) {
	// Check if token is a bearer token
	if !strings.HasPrefix(c.Request.Header.Get("Authorization"), "Bearer ") {
		panic(errors.New("invalid token"))
	}

	// Load secret key
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	panic(err)
	// 	log.Fatal("Error loading .env file")
	// }
	// mySigningKey := os.Getenv("SECRET_KEY")
	mySigningKey := []byte("Somethingveryimportantmbidk")

	// Validate token
	tokenString := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
