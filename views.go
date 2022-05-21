package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

func createMovie(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	// decode the request
	_ = json.NewDecoder(request.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}

func getMovies(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	for _, item := range movies {
		if params["id"] == item.ID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func updateMovie(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	params := mux.Vars(request)

	for index, item := range movies {
		if params["id"] == item.ID {
			movie.ID = item.ID
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, movie)
		}
	}
	json.NewEncoder(w).Encode(movies)

}
