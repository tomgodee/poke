package todomodel

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type Todo struct {
	ID     int    `json:"id"`
	Body   string `json:"body" form:"body"`
	Done   bool   `json:"done" form:"done"`
	UserID int    `json:"userID,omitempty" form:"userID"`
}

func GetAll(db *sql.DB) (t []Todo) {
	const query = `
	SELECT id, body, done
	FROM todos`

	var todoList []Todo
	rows, err := db.Query(query)
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
	a, err := json.Marshal(todoList)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(a))
	return todoList
}

func Create(db *sql.DB, body string, user_id int) {
	const query = `
	INSERT INTO todos (body, user_id)
	VALUES ($1, $2)
	`

	_, err := db.Exec(query, body, user_id)
	if err != nil {
		panic(err)
	}
}
