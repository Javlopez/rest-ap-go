package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have visited woooooo %s !", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handle)
	fmt.Println("Starting Web app on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

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
