package model

type User struct {
	Id 	string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	LastUpdate string `json:"last-update"`
}