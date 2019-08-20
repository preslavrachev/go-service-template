package main

import (
	"github.com/go-chi/chi"
	"github.com/preslavrachev/go-service-template/config"
	"github.com/preslavrachev/go-service-template/config/webservice"
	http2 "github.com/preslavrachev/go-service-template/http"
	"github.com/preslavrachev/go-service-template/persistence"
	"log"
	"net/http"
)

func main() {

	db := &persistence.DB{DB: webservice.SetUpDB()}
	appContext := &config.AppContext{
		Server: http2.NewServer(chi.NewRouter()),
		DB:     db,
	}
	webservice.SetUpRoutes(appContext)

	log.Fatal(http.ListenAndServe(":8080", appContext.Server.Router()))
}
