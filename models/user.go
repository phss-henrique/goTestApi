package models

type User struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
	Numero int   `json:"numero"`
}