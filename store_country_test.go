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

func showMsg(name string, fail bool, t *testing.T) {
	if fail {
		t.Error(name + " failed.")
	} else {
		t.Log(name + " successful.")
	}
}

func showMsgWrapper(name string, err error, t *testing.T) {
	showMsg(name, err != nil, t)
}

func Test_ConnectDB(t *testing.T) {
	_, err := ConnectDB()
	showMsgWrapper(t.Name(), err, t)
}

func Test_DeleteCountries(t *testing.T) {
	err := DeleteCountries()
	showMsgWrapper(t.Name(), err, t)
}

func Test_CreateCountry(t *testing.T) {
	_, err := CreateCountry(countries[0])
	showMsgWrapper(t.Name(), err, t)
}

func Test_GetCountry(t *testing.T) {
	country, err := GetCountry(countries[0].Code)
	showMsg(t.Name(), err != nil || country.Name != countries[0].Name, t)
}

func Test_GetCountries(t *testing.T) {
	_, err := CreateCountry(countries[1])
	if err != nil {
		showMsg(t.Name(), true, t)
	}

	var cntries []Country
	cntries, err = GetCountries()

	if err != nil {
		showMsg(t.Name(), true, t)
	}
	if cntries[0].Code != countries[0].Code || cntries[0].Name != countries[0].Name ||
		cntries[1].Code != countries[1].Code || cntries[1].Name != countries[1].Name {
		showMsg(t.Name(), true, t)
	}

	showMsg(t.Name(), false, t)
}

func Test_UpdateCountry(t *testing.T) {
	var changedCountry Country
	changedCountry.Code = countries[0].Code
	changedCountry.Name = "Hindustan"

	err := UpdateCountry(changedCountry)
	if err != nil {
		showMsg(t.Name(), true, t)
	}

	var country Country
	country, err = GetCountry(changedCountry.Code)

	if err != nil {
		showMsg(t.Name(), true, t)
	}

	if country.Name != "Hindustan" {
		showMsg(t.Name(), true, t)
	}

	changedCountry.Name = countries[0].Name
	err = UpdateCountry(changedCountry)
	if err != nil {
		showMsg(t.Name(), true, t)
	}

	country, err = GetCountry(changedCountry.Code)

	if err != nil {
		showMsg(t.Name(), true, t)
	}

	if country.Name != countries[0].Name {
		showMsg(t.Name(), true, t)
	}

	showMsg(t.Name(), false, t)
}
