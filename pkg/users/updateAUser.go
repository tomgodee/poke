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

// 	updateAUser(db)
// 	defer db.Close()

// }

// func updateAUser(db *sql.DB) {
// 	sqlStatement := `
// 		UPDATE users
// 		SET first_name = $2, last_name = $3
// 		WHERE id = $1;`

// 	_, err := db.Exec(sqlStatement, 5, "Diego", "Fiorillo")
// 	if err != nil {
// 		panic(err)
// 	}
// }
