package repository

import (
	"log"
	"time"

	"github.com/axpensive/golang-todo-app-clean-architecture/app/entity"
	// "github.com/axpensive/golang-todo-app-clean-architecture/app/entity"
)

// type repoUser struct {
// 	entity.User
// }

func CreateTodo(content string, userID int) (err error) {
	cmd := `insert into todos (content,user_id, created_at) values ($1,$2,$3)`

	_, err = Db.Exec(cmd, content, userID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetTodo(todoID int) (todo entity.Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where id = $1`

	todo = []entity.Todo{}
	// Scanで出力を変数に代入します。
	err = Db.QueryRow(cmd, todoID).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)
	return todo, err
}

func GetTodosByUser(userID int) (todos []entity.Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where user_id = $1`

	rows, err := Db.Query(cmd, userID)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo entity.Todo

		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err

}

func UpdateTodo(todoID int, content string, userID int) error {
	cmd := `update todos set content = $1, user_id = $2 where id = $3`
	_, err = Db.Exec(cmd, content, userID, todoID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func DeleteTodo(todoID int) error {
	cmd := `delete from todos where id = $1`
	_, err = Db.Exec(cmd, todoID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
