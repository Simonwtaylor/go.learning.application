package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryStore())

	if err := http.ListenAndServe(":5000", server.Handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
