package location

import (
	"testing"
)

var countries = []Country{
	Country{
		Code: "IN",
		Name: "India",
	},
	Country{
		Code: "SG",
		Name: "Singapore",
	},
}

func report1(name string, fail bool, t *testing.T) {
	if fail {
		t.Error(name + " failed.")
	} else {
		t.Log(name + " successful.")
	}
}

func report(name string, err error, t *testing.T) {
	report1(name, err != nil, t)
}

func Test_ConnectDB(t *testing.T) {
	_, err := ConnectDB()
	report(t.Name(), err, t)
}

func Test_DeleteCountries(t *testing.T) {
	err := DeleteCountries()
	report(t.Name(), err, t)
}

func Test_CreateCountry(t *testing.T) {
	_, err := CreateCountry(countries[0])
	report(t.Name(), err, t)
}

func Test_GetCountry(t *testing.T) {
	country, err := GetCountry(countries[0].Code)
	report1(t.Name(), err != nil || country.Name != countries[0].Name, t)
}
