package models

type Address struct {
	Name    string `json:"-"`
	Street  string `json:"primary_line"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip_code"`
	Country string `json:"-"`
}
