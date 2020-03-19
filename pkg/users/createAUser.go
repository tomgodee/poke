package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "zxc321"
	Dbname   = "poke_development"
)

func createAUser() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	insertRecordToUsers(db)
	defer db.Close()
}

func insertRecordToUsers(db *sql.DB) {
	const query = `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	id := 0
	err := db.QueryRow(query, 3, "Fiorillo@mail.com", "Anas", "Fiorillo").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
