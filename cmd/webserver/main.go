package main

import (
	"log"
	"net/http"
	"os"

	"go.learning.application/poker"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating store %v", err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server.Handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}