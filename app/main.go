package main

import (
	lc "github.com/nkumar15/location"
	"log"
	"net/http"
)

func ServeWeb() {
	router := lc.NewRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	//CreateCountriesTable()
	//InsertCountries()
	//CreateBranch()
	ServeWeb()
	//GetCountries()
	//DeleteCountry("AU")
	//GetCountries()
	//GetCountry("IN")
}
