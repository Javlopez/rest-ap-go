package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID     int    `bson:"_id"`
	Title  string `json:"title"`
	Image  string `json:"image"`
	Price  uint64 `json:"price"`
	Rating uint8  `json:"rating"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have visited  %s !", r.URL.Path)
}

func handleGetClient(w http.ResponseWriter, r *http.Request) {

	product := Product{1234, "This is the title", "This is the image", 98, 2}

	p, err := json.Marshal(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(p)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handle).Methods("GET")
	r.HandleFunc("/api/v1/getClient/", handleGetClient).Methods("GET")

	fmt.Println("Starting Web app on 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

/*
api/v1/getClient
*/
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "No implemented tes never")
// }

// func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "not implemented yet !")
// }

// func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "not implemented yet !")
// }

// func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "not implemented yet !")
// }

// func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "not implemented yet !")
// }

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
// 	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
// 	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
// 	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
// 	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")

// 	if err := http.ListenAndServe(":8080", r); err != nil {
// 		log.Fatal(err)
// 	}

// }
