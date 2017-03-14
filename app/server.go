package main

import (
	"log"
	"net/http"
)

func main() {
	router := getRouters()
	log.Fatal(http.ListenAndServe(":8081", router))
}
