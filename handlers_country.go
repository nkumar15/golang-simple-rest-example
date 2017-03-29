package location

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetCountriesHandler(w http.ResponseWriter, r *http.Request) {

	countries, err := GetCountries()

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err = json.NewEncoder(w).Encode(countries); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func GetCountryHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	code := vars["code"]

	var country Country
	country, err := GetCountry(code)

	if err != nil {
		if err == ErrNoMoreRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(country); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CreateCountryHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var country Country
	if err := json.Unmarshal(body, &country); err != nil {
		w.WriteHeader(422) // unprocessable entity
		return
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c, err := CreateCountry(country)

	if err != nil {
		log.Println("Some error from create country")
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	} else {
		log.Println("creating proper response")
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(c); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

/*
func UpdateCityByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	var city City
	if err := json.Unmarshal(body, &city); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	city.Id = id
	c, err := UpdateCityById(city)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	if err := json.NewEncoder(w).Encode(c); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
*/

func DeleteCountryHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	vars := mux.Vars(r)
	code := vars["code"]

	err := DeleteCountry(code)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)

	}

	w.WriteHeader(http.StatusNoContent)
}
