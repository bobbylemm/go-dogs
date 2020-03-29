package server

import (
	"context"
	"fmt"
	"go-dogs/config"
	"net/http"
	"go-dogs/db"
)

func HandleGetHome() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Welcome to the dog breeds API")
	}
}

func HandleGetAllDogs(config config.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db, err := db.ConnectToDB( context.Background())
		if err != nil {
			http.Error(writer, "could not connect to DB", http.StatusInternalServerError)
		}
	}
}