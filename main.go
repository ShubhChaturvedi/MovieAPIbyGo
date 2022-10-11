package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getmovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func deletemovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)

}
func getmovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}

	}

}

func main() {
	route := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "23456", Title: "Sooryavanshi", Director: &Director{FirstName: "Ashish", LastName: "Agrawal"}})
	movies = append(movies, Movie{ID: "2", Isbn: "23766", Title: "lost in space", Director: &Director{FirstName: "Shubh", LastName: "Chaturvedi"}})
	route.HandleFunc("/movie", getmovies).Methods("GET")
	route.HandleFunc("/movie/{id}", getmovie).Methods("GET")
	route.HandleFunc("/movie", createmovie).Methods("POST")
	route.HandleFunc("/movie/{id}", updatemovie).Methods("PUT")
	route.HandleFunc("/movie/{id}", deletemovie).Methods("DELETE")

	fmt.Println("Starting golang Web Server at :3000 port")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}

}
