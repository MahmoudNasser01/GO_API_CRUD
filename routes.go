package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func LoadRoutes() {
	request := mux.NewRouter()
	request.HandleFunc("/movies", getMovies).Methods("GET")
	request.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	request.HandleFunc("/movies", createMovie).Methods("POST")
	request.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	request.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// starting the server message
	fmt.Println("starting the server at port 8000!")

	log.Fatal(http.ListenAndServe(":8000", request))

}
