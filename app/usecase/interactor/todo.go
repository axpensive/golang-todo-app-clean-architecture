package interactor

import (
	"github.com/axpensive/golang-todo-app-clean-architecture/app/adapter/gateway/repository"
	"github.com/axpensive/golang-todo-app-clean-architecture/app/entity"
)

func SaveNewTodo(content string, userID int) error {
	if err := repository.CreateTodo(content, userID); err != nil {
		return err
	}
	return nil
}

func GetTodo(userID int) (entity.Todo, error) {
	todos, err := repository.GetTodosByUser(userID)
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func GetTodosByUserID(userID int) (entity.Todo, error) {
	todos, err := repository.GetTodosByUser(userID)
	if err != nil {
		return todos, err
	}
	return todos, err
}

func UpdateTodo(todoID int, content string, userID int) error {
	if err := repository.UpdateTodo(todoID, content, userID); err != nil {
		return err
	}

	return nil
}

func DeleteTodo(todoID int) error {
	if err := repository.DeleteTodo(todoID); err != nil {
		return err
	}

	return nil
}
