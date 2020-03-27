package db

import (
	"io/ioutil"
	"log"
	"net/http"
)

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
