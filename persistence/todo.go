package persistence

import (
	"github.com/jinzhu/gorm"
	myservice "github.com/preslavrachev/go-service-template"
)

type Todo struct {
	gorm.Model
	Name string
}

type DB struct {
	*gorm.DB
}

func (db *DB) FindAllTodos() []*myservice.Todo {
	var todos []*Todo
	db.Find(&todos)

	var results []*myservice.Todo
	for _, todo := range todos {
		results = append(results, &myservice.Todo{
			Name: todo.Name,
		})
	}

	return results
}

func (db *DB) SaveTodo(t *myservice.Todo) {
	todoEntity := &Todo{
		Name: t.Name,
	}

	db.Save(todoEntity)
}
