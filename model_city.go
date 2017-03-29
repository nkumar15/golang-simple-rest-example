package location

type City struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Cities []City
