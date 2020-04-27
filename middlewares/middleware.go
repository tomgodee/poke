package middlewares

import (
	"errors"
	"fmt"
	"os"
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
	mySigningKey := []byte(os.Getenv("SECRET_KEY"))

	// Validate token
	tokenString := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err == jwt.ErrSignatureInvalid || !token.Valid {
		c.JSON(401, "Unauthorized")
	}

	if err != nil {
		panic(err)
	}
}
