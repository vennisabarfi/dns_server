package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Println("Error starting server", err)
	}

}
