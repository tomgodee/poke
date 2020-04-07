package userModel

import (
	"database/sql"
)

type User struct {
	ID       int    `json: "id"`
	Username string `json: "username"`
	Password string `json: "password`
	Email    string `json: "email"`
}

func GetAUser(db *sql.DB, user_id int) (u User) {
	const query = `
	SELECT id, username, email FROM users
	WHERE id = $1;`

	// var id int
	// var username, email string
	var user User

	row := db.QueryRow(query, user_id)

	err := row.Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		panic(err)
	}
	return user
}
