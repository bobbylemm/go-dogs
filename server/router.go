package server

import (
	"github.com/gorilla/mux"
	"go-dogs/config"
)

type Router struct {
	*mux.Router
	AppConfig config.Config
}

func (r Router) initializeRoutes() {
	r.HandleFunc("/", HandleGetHome())
	r.HandleFunc("/dogs", HandleGetAllDogs(r.AppConfig))
}

func NewRouter(appConfig config.Config) *Router {
	return &Router{mux.NewRouter(), appConfig}
}