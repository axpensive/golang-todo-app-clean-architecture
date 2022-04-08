package port

import (
	"github.com/axpensive/golang-todo-app-clean-architecture/app/entity"
)

type IUserInputPort interface {
	SaveTodo(content string, user *entity.Todo) error
}

type IUserOutputPort interface {
	OutputTodo(*entity.User)
	OutputTodos([]*entity.User)
	OutputError(error) error
}

type IUserRepository interface {
	CreateUser(name string, email string, password string) (err error)
	Getuser(id int) (user entity.User, err error)
	UpdateUser(name string, email string, id int) (err error)
	DeleteUser(id int) (err error)
	GetUserByEmail(email string) (user entity.User, err error)
}
