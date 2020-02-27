package main

import (
	"fmt"
	"log"
	"net/http"

	"go.learning.application/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server.Handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
