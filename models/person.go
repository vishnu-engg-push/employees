package models

type Person struct {
	Id      int     `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Active  bool    `json:"active,omitempty"`
	Address Address `json:"address,omitempty"`
}
