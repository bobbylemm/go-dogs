package services

import (
	"encoding/json"
	"go-dogs/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"net/http"
)

func (d *DB) GetAllDogs(r *http.Request) ([]models.Dog, error) {
	var dog models.Dog
	var result []models.Dog
	cur, err := d.Collection("dogs").Find(r.Context(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(r.Context())
	for cur.Next(r.Context()) {
		err := cur.Decode(&dog)
		result = append(result, dog)
		if err != nil { log.Fatal(err) }
	}
	return result, nil
}

func (d *DB) AllBreeds(r *http.Request) ([]models.Breed,error) {
	resp, err := http.Get("https://api.thedogapi.com/v1/breeds")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var allBreeds []models.Breed
	var dataFromDB models.Breed

	err = json.Unmarshal(body, &allBreeds)
	if err != nil {
		log.Fatal(err)
	}
	cur, err := d.Collection("breeds").Find(r.Context(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(r.Context())
	for cur.Next(r.Context()) {
		err := cur.Decode(&dataFromDB)
		allBreeds = append(allBreeds, dataFromDB)
		if err != nil { log.Fatal(err) }
	}
	return allBreeds, nil
}

func (d *DB) AddDog(r *http.Request) (*mongo.InsertOneResult, error) {
	var dog models.Dog
	if err := json.NewDecoder(r.Body).Decode(&dog); err != nil {
		return nil, err
	}
	response, err := d.Collection("dogs").InsertOne(r.Context(), dog)
	if err != nil {
		return nil, err
	}
	return response, nil
}