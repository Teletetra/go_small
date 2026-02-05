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
	Isbn     string    `json:"Isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// this is a slice
var movies []Movie

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&movie)
	if err != nil {
		http.Error(w, "bad status", http.StatusBadRequest)
		return
	}
	fmt.Println("🔥 createMovie HIT, method =", r.Method)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

//	func getMovie(w http.ResponseWriter,r *http.Request){
//		w.Header().Set("Content-Type","application/json")
//		params:=mux.Vars(r)
//		for _,item:=range movies{
//			if item.ID==params["id"]{
//				json.NewEncoder(w).Encode(item)
//				return
//			}
//		}
//	}
func getMovieB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for _, movie := range movies {
		if movie.ID == id {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

// func createMovie(w http.ResponseWriter,r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	var movie Movie
// 	_=json.NewDecoder(r.Body).Decode(&movie)
// 	movie.ID=strconv.Itoa(rand.Intn(1000000000))
// 	movies=append(movies, movie)
// 	json.NewEncoder(w).Encode(movie)
// }

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			// _ = json.NewDecoder(r.Body).Decode(&movie)
			// movie.ID = params["id"]
			// movies = append(movies, movie)
			movies[index]=movie
			movies[index].ID=params["id"]
			json.NewEncoder(w).Encode(movies[index])
		}
	}
}

// func enableCORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 		// 👇 preflight ko yahin nipta do
// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})}
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.Use(corsMiddleware)
	movies = append(movies, Movie{ID: "1", Isbn: "34223", Title: "movie one", Director: &Director{Firstname: "john", Lastname: "smith"}})
// r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "pong")
// }).Methods("GET")
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("/movies", getAllMovies).Methods("GET")

	r.HandleFunc("/movies/{id}", getMovieB).Methods("GET")

	r.HandleFunc("/createMovie", createMovie).Methods("POST")

	r.HandleFunc("/updateMovie/{id}", updateMovie).Methods("PUT")

	r.HandleFunc("/deleteMovie/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting at server 8083\n")
log.Fatal(http.ListenAndServe(":8083", r))
}	
