package config

import (
	"github.com/preslavrachev/go-service-template/http"
	"github.com/preslavrachev/go-service-template/persistence"
)

type AppContext struct {
	Server *http.Server
	DB     *persistence.DB
}
