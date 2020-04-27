package usermodel

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json: "id"`
	Username string `json: "username" form:"username"`
	Password string `json: "password" form:"password"`
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

	hash := hashPwd(data["password"])

	var lastInsertID int
	err := db.QueryRow(query, data["username"], hash, data["email"]).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}

	return lastInsertID
}

func Update(db *sql.DB, data map[string]string, id int) {
	const query = `
	UPDATE users
	SET Username = $1
	WHERE ID = $2
	RETURNING ID
	`

	_, err := db.Exec(query, data["username"], id)
	if err != nil {
		panic(err)
	}

	// TODO: Somehow res is the address in memory ???
	// => Need to further research to differentiate db.Exec vs db.Query
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
	SELECT password, id
	FROM users
	WHERE username = $1`

	var pwd string
	var id int
	row := db.QueryRow(query, loginData["username"])
	err = row.Scan(&pwd, &id)
	switch {
	case err == sql.ErrNoRows:
		return err
	case err != nil:
		panic(err)
	}

	// Check if password is correct
	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(loginData["password"]))
	switch {
	case err == bcrypt.ErrMismatchedHashAndPassword:
		return err
	case err != nil:
		panic(err)
	}
	createToken(string(id))

	return nil
}

func createToken(userID string) {
	// Load secret key
	mySigningKey := []byte(os.Getenv("SECRET_KEY"))

	// Set token's expiration time
	expirationTime := time.Now().Add(60 * 24 * time.Minute)
	claims := &jwt.StandardClaims{
		Audience:  userID,
		ExpiresAt: expirationTime.Unix(),
	}

	// Create and sign the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println(ss)
	fmt.Println(err)
}

func hashPwd(pwd string) (hash []byte) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return hash
}
