package location

import (
	"net/http"
)

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

func (env *Env) cityRoutes() Routes {
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
	return cityRoutes
}

func (env *Env) countryRoutes() Routes {
	var countryRoutes = Routes{
		Route{
			"GetCountriesHandler",
			"GET",
			"/countries",
			env.GetCountriesHandler,
		},
		Route{
			"GetCountryHandler",
			"GET",
			"/countries/{code}",
			env.GetCountryHandler,
		},
		Route{
			"CreateCountryHandler",
			"POST",
			"/countries",
			env.CreateCountryHandler,
		},
		Route{
			"DeleteCountryHandler",
			"DELETE",
			"/countries/{code}",
			env.DeleteCountryHandler,
		},
		//Route{
		//	"UpdateCountryHandler",
		//	"PUT",
		//	"/countries/{id}",
		//	UpdateCityByIdHandler,
		//},
	}
	return countryRoutes
}
