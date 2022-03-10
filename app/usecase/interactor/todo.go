package interactor

type Todo struct {
	OutputPort port.TodoOutputPort
	TodoRepo   port.TodoRepository
}

func (t *Todo) GetTodoByUserID() {

}
