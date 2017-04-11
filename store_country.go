package location

import (
	"log"
	"time"

	"upper.io/db.v3"
)

func logIfError(err error) {

	if err != nil {
		log.Println("Error: ", err)
	}

}

//GetCountries ...
func GetCountries() ([]Country, error) {

	sess, err := ConnectDB()
	logIfError(err)

	defer sess.Close()

	countries := make([]Country, 0)

	col := sess.Collection("Countries")
	res := col.Find()
	defer res.Close()

	err = res.All(&countries)
	if err == db.ErrNoMoreRows {
		return make([]Country, 0), nil
	}

	return countries, err
}

//GetCountry ...
func GetCountry(code string) (Country, error) {

	sess, err := ConnectDB()
	logIfError(err)

	defer sess.Close()

	var country Country

	col := sess.Collection("Countries")
	res := col.Find(db.Cond{"Code": code})
	defer res.Close()

	err = res.One(&country)
	if err == db.ErrNoMoreRows {
		return country, ErrNoMoreRows
	}

	return country, err
}

//CreateCountry ...
func CreateCountry(c Country) (Country, error) {

	sess, err := ConnectDB()
	logIfError(err)

	defer sess.Close()

	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	c.DeletedAt = nil

	id, err := sess.Collection("Countries").Insert(c)
	if err != nil {
		return Country{}, err
	}

	if i, ok := id.(int); ok {
		c.ID = i
	}

	return c, err
}

//DeleteCountry ...
func DeleteCountry(code string) error {

	sess, err := ConnectDB()
	logIfError(err)

	defer sess.Close()

	col := sess.Collection("Countries")

	res := col.Find(db.Cond{"code": code})
	defer res.Close()
	err = res.Delete()

	return err
}

//DeleteCountries ...
func DeleteCountries() error {

	sess, err := ConnectDB()
	logIfError(err)

	defer sess.Close()

	col := sess.Collection("Countries")

	res := col.Find()
	defer res.Close()
	err = res.Delete()

	return err
}

//UpdateCountry ...
func UpdateCountry(country Country) error {
	sess, err := ConnectDB()
	logIfError(err)

	defer sess.Close()

	col := sess.Collection("Countries")

	res := col.Find("Code", country.Code)
	defer res.Close()

	err = res.Update(Country{
		Name: country.Name,
		CommonFields: CommonFields{
			UpdatedAt: time.Now(),
		},
	})

	return err

}
