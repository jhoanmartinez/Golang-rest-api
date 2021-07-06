package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	id       string    `json: "id" `
	isbn     string    `json: "isbn" `
	title    string    `json: "title" `
	director *Director `json: "director" `
}

type Director struct {
	firstname string `json: "firstname" `
	lastname  string `json: "lastname" `
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.id == params["id"] {
			fmt.Println("endpoint: /movies/{id}")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new_movie Movie
	_ = json.NewDecoder(r.Body).Decode(&new_movie)
	new_movie.id = strconv.Itoa(rand.Intn(100))
	movies = append(movies, new_movie)
	json.NewEncoder(w).Encode(new_movie)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new_movie Movie
	_ = json.NewDecoder(r.Body).Decode(&new_movie)
	new_movie.id = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, new_movie)
	json.NewEncoder(w).Encode(new_movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var old_movie Movie
			_ = json.NewDecoder(r.Body).Decode(&old_movie)
			old_movie.id = params["id"]
			movies = append(movies, old_movie)
			json.NewEncoder(w).Encode(old_movie)
			return
		}
	}
}

func main() {
	//data simulada en la base de datos
	movies = append(movies, Movie{id: "1", isbn: "111QEW", title: "1 The title 1", director: &Director{firstname: "A firt name 1", lastname: "Y last name 1"}})
	movies = append(movies, Movie{id: "2", isbn: "222QEW", title: "2 The title ", director: &Director{firstname: "B firt name 2", lastname: "Z last name 2"}})

	//Router urls con metodo de parametro
	r := mux.NewRouter()
	r.HendleFunc("/movies", getMovies).Methods("GET")
	r.Handlefunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.Handlefunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
}
