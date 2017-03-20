package main

import (
	"log"
	"time"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

type Country struct {
	ID   int    `db:"Id,omitempty"`
	Code string `db:"Code,omitempty"`
	Name string `db:"Name,omitempty"`
	CommonFields
}

func logIfError(err error) {

	if err != nil {
		log.Println(err)
	}

}

var settings = sqlite.ConnectionURL{
	Database: `./db/migrations/database.sqlite`,
}

func openDb() sqlbuilder.Database {
	sess, err := sqlite.Open(settings)

	if err != nil {
		log.Fatal("sqlite.open: %s", err)
		panic(err)
	}

	return sess
}

func GetCountries() ([]Country, error) {

	sess := openDb()
	defer sess.Close()

	countries := make([]Country, 0)

	col := sess.Collection("Countries")
	res := col.Find()

	err := res.All(&countries)

	if err == db.ErrNoMoreRows {
		return make([]Country, 0), nil
	}
	return countries, err
}

func GetCountry(code string) (Country, error) {

	sess := openDb()
	defer sess.Close()

	var country Country

	col := sess.Collection("Countries")
	res := col.Find(db.Cond{"Code": code})

	err := res.One(&country)

	if err == db.ErrNoMoreRows {
		return country, ErrNoMoreRows
	}

	return country, err
}

func CreateCountry(c Country) (Country, error) {
	sess := openDb()
	defer sess.Close()

	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	c.DeletedAt = nil

	id, err := sess.Collection("Countries").Insert(c)
	logIfError(err)
	log.Println(id)
	return c, err
}

func DeleteCountry(code string) error {

	sess := openDb()
	defer sess.Close()

	col := sess.Collection("Countries")
	res := col.Find(db.Cond{"code": code})
	err := res.Delete()

	return err
}
