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
	CreateCountriesTable()
	InsertCountries()
	ServeWeb()
	//GetCountries()
	//DeleteCountry("AU")
	//GetCountries()
	//GetCountry("IN")
}
