package http

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/preslavrachev/go-service-template"
	"net/http"
)

type Server struct {
	router *chi.Mux
}

func (s *Server) Router() *chi.Mux {
	return s.router
}

func NewServer(router *chi.Mux) *Server {
	return &Server{router: router}
}

func (s *Server) HandleTodoListGet(todoHandler interface {
	FindAllTodos() []*myservice.Todo
}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(todoHandler.FindAllTodos())
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}
}

func (s *Server) HandleTodoPost(handler interface {
	SaveTodo(t *myservice.Todo)
}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := &myservice.Todo{}
		err := json.NewDecoder(r.Body).Decode(todo)
		if err != nil {
			http.Error(w, http.StatusText(422), 422)
			return
		}

		handler.SaveTodo(todo)
	}
}
