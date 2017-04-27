package location

import (
	"errors"
	"testing"

	"upper.io/db.v3/sqlite"
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
		t.Log(t.Name() + " success.")
	}
}

func TestCountry(t *testing.T) {

	var sqliteSettings = sqlite.ConnectionURL{
		Database: `D:\programming\database\location\database.sqlite`,
	}
	db, err := sqlite.Open(sqliteSettings)
	if err != nil {
		t.Error(t.Name() + " connection failed" + err.Error())
	}

	var env Env
	env.Database.DB = db
	defer env.Database.DB.Close()

	t.Run("DeleteCountries", func(t *testing.T) { CountryDeleteAllTest(t, env) })
	t.Run("CreateCountry", func(t *testing.T) { CountryCreateTest(t, env) })
	t.Run("GetCountry", func(t *testing.T) { CountryGetTest(t, env) })
	t.Run("GetCountries", func(t *testing.T) { CountryGetAllTest(t, env) })
	t.Run("UpdateCountry", func(t *testing.T) { CountryUpdateTest(t, env) })
	t.Run("DeleteCountry", func(t *testing.T) { CountryDeleteTest(t, env) })
	t.Run("DeleteCountries", func(t *testing.T) { CountryDeleteAllTest(t, env) })
}

func CountryCreateTest(t *testing.T, env Env) {

	_, err := env.Database.CreateCountry(countries[0])
	showMsg(err, t)
}

func CountryGetTest(t *testing.T, env Env) {

	country, err := env.Database.GetCountry(countries[0].Code)
	if err != nil {
		showMsg(err, t)
	} else if country.Name != countries[0].Name {
		showMsg(errors.New("Country name mismatch"), t)
	}
}

func CountryGetAllTest(t *testing.T, env Env) {

	_, err := env.Database.CreateCountry(countries[1])
	showMsg(err, t)

	var cntries []Country
	cntries, err = env.Database.GetCountries()

	showMsg(err, t)
	if cntries[0].Code != countries[0].Code || cntries[0].Name != countries[0].Name ||
		cntries[1].Code != countries[1].Code || cntries[1].Name != countries[1].Name {
		showMsg(errors.New("Country name code mismatch"), t)
	}
	showMsg(nil, t)
}

func CountryUpdateTest(t *testing.T, env Env) {

	var changedCountry Country
	changedCountry.Code = countries[0].Code
	changedCountry.Name = "Hindustan"

	err := env.Database.UpdateCountry(changedCountry)
	showMsg(err, t)

	var country Country
	country, err = env.Database.GetCountry(changedCountry.Code)

	showMsg(err, t)

	if country.Name != "Hindustan" {
		showMsg(errors.New("Country name mismatch"), t)
	}

	changedCountry.Name = countries[0].Name
	err = env.Database.UpdateCountry(changedCountry)
	showMsg(err, t)

	country, err = env.Database.GetCountry(changedCountry.Code)

	showMsg(err, t)

	if country.Name != countries[0].Name {
		showMsg(errors.New("Country name mismatch"), t)
	}

	showMsg(nil, t)
}

func CountryDeleteTest(t *testing.T, env Env) {
	code := countries[1].Code
	err := env.Database.DeleteCountry(code)
	showMsg(err, t)
}

func CountryDeleteAllTest(t *testing.T, env Env) {

	err := env.Database.DeleteCountries()
	showMsg(err, t)
}
