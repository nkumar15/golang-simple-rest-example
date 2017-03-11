package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Country struct {
	gorm.Model
	Code string `sql:"char(2);unique"`
	Name string `sql:varchar(100);unique"`
}

func (c *Country) TableName() string {
	return "Countries"
}

func CreateCountryTable() {
	db, err := gorm.Open("sqlite3", "./logistics.db")

	if err != nil {
	}

	if !db.HasTable(&Country{}) {
		db.CreateTable(&Country{})
	}

}

func CreateCountry(c Country) {
	db, err := gorm.Open("sqlite3", "./logistics.db")

	if err != nil {
	}
	defer db.Close()

	db.NewRecord(c)
	db.Create(&c)
}
