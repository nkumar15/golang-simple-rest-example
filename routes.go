package location

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

var cityRoutes = Routes{
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

var countryRoutes = Routes{
	Route{
		"GetCountriesHandler",
		"GET",
		"/countries",
		GetCountriesHandler,
	},
	Route{
		"GetCountryHandler",
		"GET",
		"/countries/{code}",
		GetCountryHandler,
	},
	Route{
		"CreateCountryHandler",
		"POST",
		"/countries",
		CreateCountryHandler,
	},
	Route{
		"DeleteCountryHandler",
		"DELETE",
		"/countries/{code}",
		DeleteCountryHandler,
	},
	//Route{
	//	"UpdateCountryHandler",
	//	"PUT",
	//	"/countries/{id}",
	//	UpdateCityByIdHandler,
	//},
}

var branchRoutes = Routes{
	Route{
		"GetBranchesHandler",
		"GET",
		"/branches",
		GetBranchesHandler,
	},
	Route{
		"CreateBrancheHandler",
		"POST",
		"/branches",
		CreateBranchHandler,
	},
}
