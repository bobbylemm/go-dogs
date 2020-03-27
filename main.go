package main

import (
	"context"
	"github.com/gorilla/mux"
	"go-dogs/config"
	"go-dogs/controllers"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

func main() {
	c := config.DbConnect()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("db Connected!")
	}
	r := mux.NewRouter()
	r.HandleFunc("/breeds", controllers.AllBreeds).Methods(http.MethodGet)
	r.HandleFunc("/search/breeds", controllers.SearchBreed).Methods(http.MethodGet)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
