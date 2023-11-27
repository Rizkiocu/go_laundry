package model

type Employee struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Phone_number string `json:"phone_number"`
	Address      string `json:"address"`
}
