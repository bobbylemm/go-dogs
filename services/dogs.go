package services

import (
	"context"
	"go-dogs/db"
	"go-dogs/models"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
)
type DB db.DB

func (d *DB) GetAllDogs() ([]models.Dog, error) {
	var dogs []models.Dog
	allDogs, err := d.Collection("dogs").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer allDogs.Close(context.Background())
	err = allDogs.Decode(&dogs)
	if err != nil {
		return nil, err
	}
	return dogs, nil
}
func AllDogs() {
	resp, err := http.Get("https://thedogapi.com/v1/images?api_key=ce9eb0c2-8a22-4c06-8e59-898617a08303")
	if err != null {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}