package main

import (
	"fmt"
	"log"

	"github.com/VictorTarnovski/hollow-knight-api/api"
	"github.com/VictorTarnovski/hollow-knight-api/storage"
)

func main() {
	store, err := storage.NewPostgresStore()

	if err != nil {
		log.Fatal(err)

	}
	server := api.NewAPIServer("127.0.0.1:3000", store)
	fmt.Println("Booting server...")
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
