package todomodel

import (
	"database/sql"
	"log"
)

type Todo struct {
	ID     int    `json:"id"`
	Body   string `json:"body" form:"body"`
	Done   bool   `json:"done" form:"done"`
	UserID int    `json:"userID,omitempty" form:"userID"`
}

func GetAll(db *sql.DB, userID int) (t []Todo) {
	const query = `
	SELECT id, body, done
	FROM todos
	WHERE user_id = $1`

	var todoList []Todo
	rows, err := db.Query(query, userID)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Body, &todo.Done)
		if err != nil {
			log.Fatal(err)
		}

		todoList = append(todoList, todo)
	}
	if err != nil {
		log.Fatal(err)
	}
	return todoList
}

func GetOne(db *sql.DB, todoID int) (t Todo) {
	const query = `
	SELECT id, body, done
	FROM todos
	WHERE id = $1`

	var todo Todo
	row := db.QueryRow(query, todoID)
	err := row.Scan(&todo.ID, &todo.Body, &todo.Done)
	if err != nil {
		log.Fatal(err)
	}

	return todo
}

func Create(db *sql.DB, body string, userID int) {
	const query = `
	INSERT INTO todos (body, user_id)
	VALUES ($1, $2)
	`

	_, err := db.Exec(query, body, userID)
	if err != nil {
		panic(err)
	}
}

func Update(db *sql.DB, payload map[string]string, todoID int) {
	const query = `
	UPDATE todos
	SET body = $1, done = $2
	WHERE id = $3`

	_, err := db.Exec(query, payload["body"], payload["done"], todoID)
	if err != nil {
		panic(err)
	}
}

func Delete(db *sql.DB, todoID int) {
	const query = `
	DELETE FROM todos
	WHERE id = $1`

	_, err := db.Exec(query, todoID)
	if err != nil {
		panic(err)
	}
}
