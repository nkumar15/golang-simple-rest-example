package location

import (
	"log"
	"time"
	"upper.io/db.v2"
)

func logIfError(err error) {

	if err != nil {
		log.Println("Error: ", err)
	}

}

func GetCountries() ([]Country, error) {

	sess := ConnectDB()
	defer sess.Close()

	countries := make([]Country, 0)

	col := sess.Collection("Countries")

	res := col.Find()
	defer res.Close()

	err := res.All(&countries)

	if err == db.ErrNoMoreRows {
		return make([]Country, 0), nil
	}
	return countries, err
}

func GetCountry(code string) (Country, error) {

	sess := ConnectDB()
	defer sess.Close()

	var country Country

	col := sess.Collection("Countries")

	res := col.Find(db.Cond{"Code": code})
	defer res.Close()

	err := res.One(&country)

	if err == db.ErrNoMoreRows {
		return country, ErrNoMoreRows
	}

	return country, err
}

func CreateCountry(c Country) (Country, error) {
	sess := ConnectDB()
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

	sess := ConnectDB()
	defer sess.Close()

	col := sess.Collection("Countries")

	res := col.Find(db.Cond{"code": code})
	err := res.Delete()

	return err
}
