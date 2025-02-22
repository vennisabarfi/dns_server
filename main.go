package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", HealthServer)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Println("Error starting server", err)
	}

}

func HealthServer(w http.ResponseWriter, r *http.Request) {
	log.Println("Server is healthy")
	w.Write([]byte("Server is healthy"))

}
