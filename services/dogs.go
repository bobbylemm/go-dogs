package services

import (
	"context"
	"go-dogs/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (d *DB) GetAllDogs() ([]models.Dog, error) {
	var dog models.Dog
	var result []models.Dog
	cur, err := d.Collection("dogs").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		err := cur.Decode(&dog)
		result = append(result, dog)
		if err != nil { log.Fatal(err) }
	}
	return result, nil
}
//func AllDogs() {
//	resp, err := http.Get("https://thedogapi.com/v1/images?api_key=ce9eb0c2-8a22-4c06-8e59-898617a08303")
//	if err != null {
//		log.Fatalln(err)
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	log.Println(string(body))
//}