package main

import (
	"github.com/go-chi/chi"
	"github.com/preslavrachev/go-service-template/config"
	"github.com/preslavrachev/go-service-template/config/myservice"
	http2 "github.com/preslavrachev/go-service-template/http"
	"github.com/preslavrachev/go-service-template/logic"
	"github.com/preslavrachev/go-service-template/persistence"
	"log"
	"net/http"
)

func main() {

	appContext := &config.AppContext{
		Server: http2.NewServer(chi.NewRouter()),
		DB:     &persistence.DB{DB: myservice.SetUpDB()},
	}

	appContext.Server.Router().Route("/todo", func(r chi.Router) {
		setUpTodoAllGet(r, appContext)
		setUpTodoPost(r, appContext)
	})

	log.Fatal(http.ListenAndServe(":8080", appContext.Server.Router()))
}

func setUpTodoAllGet(r chi.Router, ctx *config.AppContext) {
	todosHandler := &logic.FindAllTodosHandler{TodoHandler: ctx.DB}
	r.Get("/", ctx.Server.HandleTodoListGet(todosHandler))
}

func setUpTodoPost(r chi.Router, ctx *config.AppContext) {
	saveAction := &logic.SaveTodoHandler{TodoHandler: ctx.DB}
	r.Post("/", ctx.Server.HandleTodoPost(saveAction))
}
