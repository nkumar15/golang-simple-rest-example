package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"FindCitiesHandler",
		"GET",
		"/cities",
		FindCitiesHandler,
	},
	Route{
		"FindCityByIdHandler",
		"GET",
		"/cities/{id}",
		FindCityByIdHandler,
	},
	Route{
		"CreateCityHandler",
		"POST",
		"/cities",
		CreateCityHandler,
	},
	Route{
		"DeleteCityHandler",
		"DELETE",
		"/cities/{id}",
		DeleteCityByIdHandler,
	},
	Route{
		"UpdateCityHandler",
		"PUT",
		"/cities/{id}",
		UpdateCityByIdHandler,
	},
}
