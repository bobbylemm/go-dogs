package main

import (
	"go-dogs/server"
	"log"
)

func main() {
	err := server.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
