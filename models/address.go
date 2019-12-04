package models

type Address struct {
	Name    string `json:"-"`
	Street  string `json:"primary_line"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip_code"`
	Country string `json:"-"`
}

func (a Address) Valid() bool {
	for _, val := range []string{a.Street, a.State, a.City, a.State, a.Zip, a.Country} {
		if val == "" {
			return false
		}
	}
	return true
}
