package location

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter ...
func (env *Env) NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range env.countryRoutes() {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	for _, route := range env.cityRoutes() {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
