package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-dogs/config"
	"go-dogs/models"
	"io/ioutil"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type DB struct {
	db
}

func AllBreeds(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.thedogapi.com/v1/breeds")
	//?api_key=ce9eb0c2-8a22-4c06-8e59-898617a08303
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := models.Breed{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	respondWithJson(w, http.StatusOK, data)
}

func SearchBreed(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query.Get("q"))
	resp, respErr := http.Get("https://api.thedogapi.com/v1/breeds/search?q="+query.Get("q"))
	if respErr != nil {
		log.Fatalln(respErr)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := models.Breed{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	respondWithJson(w, http.StatusOK, data)
}

func AddDog(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var dogs models.Dog
	if err := json.NewDecoder(r.Body).Decode(&dogs); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
	}
	db := config.DbConnect().Database("dogsDb")
	collection := db.Collection("dogs")
	collection.InsertOne(context.Background(), dogs)
	respondWithJson(w, http.StatusOK, data)
}