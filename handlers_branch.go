package location

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateBranchHandler(w http.ResponseWriter, r *http.Request) {
	var branch Branch
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	if err := json.Unmarshal(body, &branch); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	b, err := CreateBranch(branch)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if err := json.NewEncoder(w).Encode(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
}

func GetBranchesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	branches, err := GetBranches()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(branches); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
