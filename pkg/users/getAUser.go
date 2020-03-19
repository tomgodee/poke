// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "zxc321"
// 	dbname   = "poke_development"
// )

// func main() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// 	getAUser(db)
// 	defer db.Close()
// }

// func getAUser(db *sql.DB) {
// 	const query = `
// 	SELECT id, first_name, last_name FROM users
// 	WHERE id = $1;`

// 	var id int
// 	var firstName, lastName string

// 	row := db.QueryRow(query, 5)

// 	err := row.Scan(&id, &firstName, &lastName)

// 	switch err {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 	case nil:
// 		fmt.Println(id, firstName, lastName)
// 	default:
// 		panic(err)
// 	}
// }
