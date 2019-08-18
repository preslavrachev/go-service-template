package logic

import (
	"github.com/preslavrachev/go-service-template"
)

type FindAllTodosHandler struct {
	TodoHandler interface {
		FindAllTodos() []*myservice.Todo
	}
}

func (p *FindAllTodosHandler) FindAllTodos() []*myservice.Todo {
	return p.TodoHandler.FindAllTodos()
}

type SaveTodoHandler struct {
	TodoHandler interface {
		SaveTodo(t *myservice.Todo)
	}
}

func (p *SaveTodoHandler) SaveTodo(t *myservice.Todo) {
	p.TodoHandler.SaveTodo(t)
}
