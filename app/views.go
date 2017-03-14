package main
import "strconv"

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving index view")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, World!")
}


func handleCategories(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "GET":
		categoryList, error := getCategories()
		if error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(res, string(categoryList))

	case "POST":
		var category Category
		fmt.Println(1, req.Body)
		decoder := json.NewDecoder(req.Body)
		error := decoder.Decode(&category)
		if error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}
		createCategory(category.Name)
	}

}


func handleCategory(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	categoryId, error := strconv.Atoi(vars["categoryId"])
	if error != nil {
		http.Error(res, error.Error(), http.StatusBadRequest)
		return 
	}

	switch req.Method {
	
	case "GET":
		category, error := getCategory(categoryId)
		if error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return 
		}
		fmt.Fprint(res, string(category))
	
	case "PUT":
		var category Category
		decoder := json.NewDecoder(req.Body)
		error := decoder.Decode(&category)
		if error != nil {
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}
		updateCategory(categoryId, category.Name)

	case "DELETE":
		deleteCategory(categoryId)
	}

}