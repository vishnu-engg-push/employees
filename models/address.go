package models

type Address struct {
	City    string `json:"city,omitempty"`
	Pincode int    `json:"pincode,omitempty"`
}
