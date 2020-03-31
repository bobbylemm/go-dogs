package server

import (
	"github.com/gorilla/mux"
	"go-dogs/config"
	"net/http"
)

type Router struct {
	*mux.Router
	AppConfig config.Config
}

func (r *Router) initializeRoutes() {
	r.HandleFunc("/", HandleGetHome())
	r.HandleFunc("/dogs", HandleGetAllDogs(r.AppConfig)).Methods(http.MethodGet)
	r.HandleFunc("/dogs", HandleAddDog(r.AppConfig)).Methods(http.MethodPost)
	r.HandleFunc("/breeds", HandleGetAllBreeds(r.AppConfig)).Methods(http.MethodGet)
}

func NewRouter(appConfig config.Config) *Router {
	return &Router{mux.NewRouter(), appConfig}
}