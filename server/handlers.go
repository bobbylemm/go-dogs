package server

import (
	"context"
	"encoding/json"
	"fmt"
	"go-dogs/config"
	"go-dogs/services"
	"net/http"
)

func HandleGetHome() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Welcome to the dog breeds API")
	}
}

func HandleGetAllDogs(config config.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db, err := services.ConnectToDB(context.Background(), config)
		if err != nil {
			http.Error(writer, "could not connect to DB", http.StatusInternalServerError)
		}
		dogs, err := db.GetAllDogs()
		if err != nil {
			http.Error(writer, "could not fetch all the dogs", http.StatusInternalServerError)
		}
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(dogs)
	}
}