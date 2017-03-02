package main

import (
	"errors"
)

var currentId int

var cities Cities

func init() {
	currentId = 0
	CreateCity(City{Code: "BLR", Name: "Bengaluru"})
	CreateCity(City{Code: "HYD", Name: "Hyderabad"})
}

func FindCityById(id int) (City, error) {
	for _, city := range cities {
		if city.Id == id {
			return city, nil
		}
	}
	return City{}, nil
}

func CreateCity(city City) (City, error) {
	currentId += 1
	city.Id = currentId
	cities = append(cities, city)
	return city, nil
}

func DeleteCityById(id int) error {
	for i, city := range cities {
		if city.Id == id {
			cities = append(cities[:i], cities[i+1:]...)
			return nil
		}
	}
	return nil
}

func UpdateCityById(city City) (City, error) {
	for i, c := range cities {

		if city.Id == c.Id {
			cities[i].Name = city.Name
			cities[i].Code = city.Code
			return cities[i], nil
		}

	}
	return City{}, errors.New("City not found with this id")
}
