package location

import (
	"errors"
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

func showMsg(err error, t *testing.T) {

	if err != nil {
		t.Error(t.Name() + " failed with error" + err.Error())
	} else {
		t.Log(t.Name() + " successful.")
	}
}

func Test_ConnectDB(t *testing.T) {

	_, err := ConnectDB()
	showMsg(err, t)
}

func Test_DeleteCountries(t *testing.T) {

	err := DeleteCountries()
	showMsg(err, t)
}

func Test_CreateCountry(t *testing.T) {

	_, err := CreateCountry(countries[0])
	showMsg(err, t)
}

func Test_GetCountry(t *testing.T) {

	country, err := GetCountry(countries[0].Code)
	if err != nil {
		showMsg(err, t)
	} else if country.Name != countries[0].Name {
		showMsg(errors.New("Country name mismatch"), t)
	}
}

func Test_GetCountries(t *testing.T) {

	_, err := CreateCountry(countries[1])
	showMsg(err, t)

	var cntries []Country
	cntries, err = GetCountries()

	showMsg(err, t)
	if cntries[0].Code != countries[0].Code || cntries[0].Name != countries[0].Name ||
		cntries[1].Code != countries[1].Code || cntries[1].Name != countries[1].Name {
		showMsg(errors.New("Country name code mismatch"), t)
	}
	showMsg(nil, t)
}

func Test_UpdateCountry(t *testing.T) {

	var changedCountry Country
	changedCountry.Code = countries[0].Code
	changedCountry.Name = "Hindustan"

	err := UpdateCountry(changedCountry)
	showMsg(err, t)

	var country Country
	country, err = GetCountry(changedCountry.Code)

	showMsg(err, t)

	if country.Name != "Hindustan" {
		showMsg(errors.New("Country name mismatch"), t)
	}

	changedCountry.Name = countries[0].Name
	err = UpdateCountry(changedCountry)
	showMsg(err, t)

	country, err = GetCountry(changedCountry.Code)

	showMsg(err, t)

	if country.Name != countries[0].Name {
		showMsg(errors.New("Country name mismatch"), t)
	}

	showMsg(nil, t)
}

func Test_DeleteCountry(t *testing.T) {
	code := countries[1].Code
	err := DeleteCountry(code)
	showMsg(err, t)
}
