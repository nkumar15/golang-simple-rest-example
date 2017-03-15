package main

import (
	"fmt"
	"time"
	"upper.io/db.v2/lib/sqlbuilder"
	"upper.io/db.v2/sqlite"
        "upper.io/db.v2"
)

type CommonFields struct {
	CreatedAt time.Time `db:"CreatedAt"`
	UpdatedAt time.Time `db:"UpdatedAt"`
	DeletedAt *time.Time `db:"DeletedAt,omitempty"`
}

type Branch struct {
	ID   int64   `db:"Id,omitempty"`
	Code string `db:"Code"`
	Name string `db:"Name"`
	CommonFields
}

func initSqliteDb() sqlbuilder.Database {
	var settings = sqlite.ConnectionURL{
		Database: `./db/db1.db`,
	}
	sess, err := sqlite.Open(settings)
	if err != nil {
		panic(err)
	}
	return sess
}

func CreateBranch(branch Branch) (Branch, error) {

	sess := initSqliteDb()
	defer sess.Close()

	branches := sess.Collection("Branches")

	branch.CreatedAt = time.Now()
	branch.UpdatedAt = time.Now()
        branch.DeletedAt = nil

	id, err := branches.Insert(branch)
	if err != nil {
		fmt.Println(err)
		return Branch{}, err

	}
	branch.ID = id.(int64)
	return branch, err
}

func GetBranches() ([]Branch, error) {
	sess := initSqliteDb()
	defer sess.Close()

	var branches []Branch
	brs := sess.Collection("Branches")
	//res := brs.Find()
	res := brs.Find(db.Cond{"DeletedAt" : nil })
	err := res.All(&branches)

	if err != nil {
		fmt.Println(err)
	}

	return branches, err

}
