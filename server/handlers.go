package server

import (
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
		db, err := services.ConnectToDB(request.Context(), config)
		if err != nil {
			http.Error(writer, "could not connect to DB", http.StatusInternalServerError)
		}
		dogs, err := db.GetAllDogs(request)
		if err != nil {
			http.Error(writer, "could not fetch all the dogs", http.StatusInternalServerError)
		}
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(dogs)
	}
}

func HandleAddDog(config config.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db, err := services.ConnectToDB(request.Context(), config)
		if err != nil {
			http.Error(writer, "could not connect to DB", http.StatusInternalServerError)
		}
		result, err := db.AddDog(request)
		if err != nil {
			http.Error(writer, "could not add new record", http.StatusInternalServerError)
		}
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(result)
	}
}

func HandleGetAllBreeds(config config.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db, err := services.ConnectToDB(request.Context(), config)
		if err != nil {
			http.Error(writer, "could not connect to DB", http.StatusInternalServerError)
		}
		result, err := db.AllBreeds()
		if err != nil {
			http.Error(writer, "could not get all breeds", http.StatusInternalServerError)
		}
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(result)
	}
}