package main

import (
	"github.com/gorilla/mux"
)

func getRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/categories", handleCategories).Methods("GET", "POST")
	router.HandleFunc("/category/{categoryId}", handleCategory).Methods("GET", "PUT", "DELETE")
	router.HandleFunc("/movies", handleMovies)
	return router
}
