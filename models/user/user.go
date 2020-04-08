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

	var user User

	row := db.QueryRow(query, user_id)

	err := row.Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		panic(err)
	}
	return user
}

func GetUsers(db *sql.DB) (u []User) {
	const query = `
	SELECT id, username, email
	FROM users`

	var usersList []User

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			panic(err)
		}

		usersList = append(usersList, user)
	}
	return usersList
}

func CreateAUser(db *sql.DB, data map[string]string) (id int) {
	const query = `
	INSERT INTO users (username, password, email)
	VALUES ($1, $2, $3)
	RETURNING id`

	var lastInsertId int
	err := db.QueryRow(query, data["username"], data["password"], data["email"]).Scan(&lastInsertId)

	if err != nil {
		panic(err)
	}

	return lastInsertId
}
