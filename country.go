package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type Country struct {
	gorm.Model
	Code string `sql:"char(2);unique"`
	Name string `sql:varchar(100);unique"`
}

func (c *Country) TableName() string {
	return "Countries"
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func initDb() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./db1.db")

	return db, err
}

func CreateCountriesTable() {
	db, err := initDb()
	defer db.Close()
	checkError(err)

	dbc := db.DropTable(&Country{})
	checkError(dbc.Error)

	dbc = db.CreateTable(&Country{})
	checkError(dbc.Error)
}

func InsertCountries() {
	countries := []Country{{Name: "India", Code: "IN"},
		{Name: "Singapore", Code: "SG"},
		{Name: "Australia", Code: "AU"},
	}

	for _, country := range countries {
		CreateCountry(country)
	}
}

func CreateCountry(c Country) (Country, error) {
	db, err := initDb()
	defer db.Close()
	checkError(err)

	dbc := db.Create(&c)
	log.Println(dbc.Value)
	checkError(dbc.Error)

	return c, dbc.Error
}

func GetCountries() ([]Country, error) {
	db, err := initDb()
	defer db.Close()
	checkError(err)

	var countries []Country
	dbc := db.Find(&countries)
	if dbc.Error != nil {
		log.Println(dbc.Error)
		return nil, dbc.Error
	}
	return countries, nil
}

func GetCountry(code string) (Country, error) {
	db, err := initDb()
	defer db.Close()
	checkError(err)

	var country Country
	dbc := db.Where("code = ?", code).Find(&country)

	if dbc.Error != nil {
		log.Println(dbc.Error)
		return Country{}, dbc.Error
	}
	return country, nil
}

func DeleteCountry(code string) error {
	db, err := initDb()
	defer db.Close()
	checkError(err)

	dbc := db.Where("code = ?", code).Delete(Country{})
	if dbc.Error != nil {
		log.Println(dbc.Error)
	}

	return dbc.Error
}
