package userModel

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json: "id"`
	Username string `json: "username" form:"username"`
	Password string `json: "password form:"password"`
	Email    string `json: "email" form:"email"`
}

func GetOne(db *sql.DB, user_id int) (u User) {
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

func GetAll(db *sql.DB) (u []User) {
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

func Create(db *sql.DB, data map[string]string) (id int) {
	const query = `
	INSERT INTO users (username, password, email)
	VALUES ($1, $2, $3)
	RETURNING id`

	// Hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	// inputPwd := []byte(strings.Join([]string{"z", "x", "c"}, ""))
	// err = bcrypt.CompareHashAndPassword(hash, inputPwd)
	// if err != nil {
	// 	panic(err)
	// }

	var lastInsertID int
	err = db.QueryRow(query, data["username"], hash, data["email"]).Scan(&lastInsertID)

	if err != nil {
		panic(err)
	}

	return lastInsertID
}

func Update(db *sql.DB, data map[string]string, id int) {
	const query = `
	UPDATE users
	SET Username = $1, Password = $2, Email = $3
	WHERE ID = $4
	RETURNING ID
	`

	fmt.Println(data)

	res, err := db.Exec(query, data["username"], data["password"], data["email"], id)

	if err != nil {
		panic(err)
	}

	// TODO: Somehow res is the address in memory ???
	// => Need to further research to differentiate db.Exec vs db.Query
	fmt.Println(res)

	// var updateId = res.id

	return
}

func Delete(db *sql.DB, id int) {
	const query = `
	DELETE FROM users
	WHERE ID = $1
	RETURNING ID
	`
	_, err := db.Exec(query, id)

	if err != nil {
		panic(err)
	}
}

func Login(db *sql.DB, loginData map[string]string) (err error) {
	const query = `
	SELECT password
	FROM users
	WHERE username = $1`

	var pwd string
	row := db.QueryRow(query, loginData["username"])
	err = row.Scan(&pwd)
	switch {
	case err == sql.ErrNoRows:
		return err
	case err != nil:
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(loginData["password"]))
	switch {
	case err == bcrypt.ErrMismatchedHashAndPassword:
		return err
	case err != nil:
		panic(err)
	}

	return nil
}
