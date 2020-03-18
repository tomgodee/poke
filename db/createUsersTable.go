package main

import (
	"database/sql"
	"fmt"

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

	createUserTable(db)

	defer db.Close()
}

func createUsersTable(db *sql.DB) {
	const query = `
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			age INT,
			first_name TEXT,
			last_name TEXT,
			email TEXT UNIQUE NOT NULL
		)`
	_, err = db.Exec(query)

	if err != nil {
		panic(err)
	}
}
