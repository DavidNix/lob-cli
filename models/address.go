package models

type Address struct {
	Name         string `json:"-"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"address_city"`
	State        string `json:"address_state"`
	Zip          string `json:"address_zip"`
	Country      string `json:"address_country"`
}
