package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(SpecServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
