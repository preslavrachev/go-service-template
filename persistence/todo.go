package persistence

import (
	"github.com/jinzhu/gorm"
	myservice "github.com/preslavrachev/go-service-template"
)

type Todo struct {
	gorm.Model
	Name string
}

type TodoStore struct {
	db IGormDB
}

func NewTodoStore(db IGormDB) *TodoStore {
	return &TodoStore{db: db}
}

func (ts *TodoStore) FindAllTodos() []*myservice.Todo {
	var todos []*Todo
	ts.db.Find(&todos)

	var results []*myservice.Todo
	for _, todo := range todos {
		results = append(results, &myservice.Todo{
			Name: todo.Name,
		})
	}

	return results
}

func (ts *TodoStore) SaveTodo(t *myservice.Todo) {
	todoEntity := &Todo{
		Name: t.Name,
	}

	ts.db.Save(todoEntity)
}
