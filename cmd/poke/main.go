package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "zxc321"
	dbname   = "poke_development"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	createUserTable(db)
	insertRecordToUsers(db, err)
	defer db.Close()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		message := "Hello " + name
		path := c.FullPath()
		// c.String(http.StatusOK, "Hello %s", name)
		c.JSON(200, gin.H{
			"message": message,
			"path":    path,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func createUserTable(db *sql.DB) {
	const query = `
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			age INT,
			first_name TEXT,
			last_name TEXT,
			email TEXT UNIQUE NOT NULL
		)`
	db.Exec(query)
}

func insertRecordToUsers(db *sql.DB, err error) {
	const query = `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	id := 0
	err = db.QueryRow(query, 25, "tom@mail.com", "Tom", "Vu").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
