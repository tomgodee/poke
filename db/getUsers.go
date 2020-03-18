package main

import (
	"database/sql"
	"fmt"

	//
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
	getUsers(db)
	defer db.Close()
}

func getUsers(db *sql.DB) {
	const query = `
	SELECT id, first_name, last_name FROM users
	LIMIT $1`

	rows, err := db.Query(query, 5)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var firstName, lastName string

		err = rows.Scan(&id, &firstName, &lastName)

		if err != nil {
			panic(err)
		}

		fmt.Println(id, firstName, lastName)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
