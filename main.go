package main

import (
	"encoding/json" // for encoding data into json while sending to postman
	"fmt"
	"log"       // for error white connecting to servers
	"math/rand" // for id of movie
	"net/http"  // for server implementation
	"strconv"   // for converting to string format

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

var movies[] Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) // sending json encapsulated Movie type data to frontend/ Postman

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// pass id from postman ( request from postman ) 
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if (item.ID == params["id"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if (item.ID == params["id"]) {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie 
			json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = item.ID 
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID : "1", Isbn : "438227", Title : "Movie one", Director : &Director{Firstname : "John", Lastname : "Maverick"}})
	movies = append(movies, Movie{ID : "2", Isbn : "439227", Title : "Movie two", Director : &Director{Firstname : "Steve", Lastname : "Baldwin"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// STARTING SERVER
	fmt.Println("Starting server at port:8000")
	log.Fatal(http.ListenAndServe(":8000", r))




}