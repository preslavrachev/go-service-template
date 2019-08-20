package webservice

import (
	"github.com/go-chi/chi"
	"github.com/preslavrachev/go-service-template/config"
	"github.com/preslavrachev/go-service-template/logic"
)

func SetUpRoutes(appContext *config.AppContext) {
	appContext.Server.Router().Route("/todo", func(r chi.Router) {
		setUpTodoAllGet(r, appContext)
		setUpTodoPost(r, appContext)
	})
}

func setUpTodoAllGet(r chi.Router, ctx *config.AppContext) {
	todosHandler := &logic.FindAllTodosHandler{TodoHandler: ctx.DB}
	r.Get("/", ctx.Server.HandleTodoListGet(todosHandler))
}

func setUpTodoPost(r chi.Router, ctx *config.AppContext) {
	saveAction := &logic.SaveTodoHandler{TodoHandler: ctx.DB}
	r.Post("/", ctx.Server.HandleTodoPost(saveAction))
}
