package main

import (
	"log"
	"net/http"
)

func ServeWeb() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	ServeWeb()
}
