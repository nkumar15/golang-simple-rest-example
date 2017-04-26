package location

import (
	"time"

	db "upper.io/db.v2"
)

//GetCountries ...
func (lcdb *lcDatabase) GetCountries() ([]Country, error) {

	countries := make([]Country, 0)

	col := lcdb.DB.Collection("Countries")
	res := col.Find()
	defer res.Close()

	err := res.All(&countries)
	if err == ErrNoMoreRows {
		return make([]Country, 0), nil
	}

	return countries, err
}

//GetCountry ...
func (lcdb *lcDatabase) GetCountry(code string) (Country, error) {

	var country Country

	col := lcdb.DB.Collection("Countries")
	res := col.Find(db.Cond{"Code": code})
	defer res.Close()

	err := res.One(&country)
	if err == ErrNoMoreRows {
		return country, ErrNoMoreRows
	}

	return country, err
}

//CreateCountry ...
func (lcdb *lcDatabase) CreateCountry(c Country) (Country, error) {

	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	c.DeletedAt = nil

	id, err := lcdb.DB.Collection("Countries").Insert(c)
	if err != nil {
		return Country{}, err
	}

	if i, ok := id.(int); ok {
		c.ID = i
	}

	return c, err
}

//DeleteCountry ...
func (lcdb *lcDatabase) DeleteCountry(code string) error {

	col := lcdb.DB.Collection("Countries")

	res := col.Find(db.Cond{"code": code})
	defer res.Close()
	err := res.Delete()

	return err
}

//DeleteCountries ...
func (lcdb *lcDatabase) DeleteCountries() error {

	col := lcdb.DB.Collection("Countries")

	res := col.Find()
	defer res.Close()
	err := res.Delete()

	return err
}

//UpdateCountry ...
func (lcdb *lcDatabase) UpdateCountry(country Country) error {

	col := lcdb.DB.Collection("Countries")

	res := col.Find("Code", country.Code)
	defer res.Close()

	err := res.Update(Country{
		Name: country.Name,
		CommonFields: CommonFields{
			UpdatedAt: time.Now(),
		},
	})

	return err
}
