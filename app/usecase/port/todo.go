package port

import (
	"github.com/axpensive/golang-todo-app-clean-architecture/app/entity"
)

type ITodoInputPort interface {
	SaveTodo(content string, user *entity.Todo) error
}

type ITodoOutputPort interface {
	OutputTodo(*entity.Todo)
	OutputTodos([]*entity.Todo)
	OutputError(error) error
}

type ITodoRepository interface {
	CreateTodo(content string, userID int) error
	GetTodo(todoID int) (todo entity.Todo, err error)
	GetTodosByUser(userID int) (todos []entity.Todo, err error)
	UpdateTodo(todoID int, content string, userID int) error
	DeleteTodo(todoID int) error
}
