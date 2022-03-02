package model

type Response struct {
	Message string   `json:"message"`
	Number  int32    `json:"number"`
	People  []Person `json:"people"`
}
