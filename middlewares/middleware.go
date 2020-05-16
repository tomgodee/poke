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

	if err == jwt.ErrSignatureInvalid || !token.Valid || claims.Audience != c.Param("id") {
		c.JSON(401, "Unauthorized")
		c.Abort()
	}

	if err != nil {
		panic(err)
	}
}

func CorsHandler(c *gin.Context) {
	var origin = "*"
	if c.Request.Header.Get("Origin") != "" {
		origin = c.Request.Header.Get("Origin")
	} else {
		origin = "http://localhost:3000"
	}

	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, origin, Cache-Control, X-Requested-With, Referer, User-Agent")
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)

	if c.Request.Method == "OPTIONS" {
		c.Status(200)
	}
}
