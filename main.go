package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/uuid", handleRequest)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(response http.ResponseWriter, request *http.Request) {
	id := uuid.New()
	fmt.Fprintf(response, id.String())
}