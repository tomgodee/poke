package userModel

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int    `json: "id"`
	Username string `json: "username"`
	Password string `json: "password`
	Email    string `json: "email"`
}

func GetAUser(db *sql.DB, user_id int) {
	const query = `
	SELECT id, username, email FROM users
	WHERE id = $1;`

	var id int
	var username, email string
	// var user User

	row := db.QueryRow(query, user_id)

	err := row.Scan(&id, &username, &email)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, username, email)
		// user = User(id, username, email)
	default:
		panic(err)
	}

}
