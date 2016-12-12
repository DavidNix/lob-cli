package parse

import (
	"github.com/davidnix/lob-cli/models"
	cli "gopkg.in/urfave/cli.v1"
)

//FromAddress returns parsed return Address
func FromAddress(c *cli.Context) models.Address {
	return models.Address{
		Name:         c.GlobalString("from-name"),
		AddressLine1: c.GlobalString("from-address"),
		AddressLine2: "",
		City:         c.GlobalString("from-city"),
		State:        c.GlobalString("from-state"),
		Zip:          c.GlobalString("from-zip"),
		Country:      c.GlobalString("from-country"),
	}
}
