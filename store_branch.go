package location

import (
	"fmt"
	"time"
	"upper.io/db.v2"
)

func CreateBranch(branch Branch) (Branch, error) {

	sess, err := ConnectDB()
	logIfError(err)
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
	sess, err := ConnectDB()
	logIfError(err)

	defer sess.Close()

	var branches []Branch
	brs := sess.Collection("Branches")
	res := brs.Find(db.Cond{"DeletedAt": nil})
	err = res.All(&branches)

	if err != nil {
		fmt.Println(err)
	}

	return branches, err

}
