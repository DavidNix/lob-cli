package parse

import (
	"github.com/davidnix/lob-cli/models"
	cli "gopkg.in/urfave/cli.v1"
)

//FromAddress returns parsed return Address
func FromAddress(c *cli.Context) models.Address {
	return models.Address{
		Name:         c.GlobalString("from-name"),
		Street:       c.GlobalString("from-address"),
		City:         c.GlobalString("from-city"),
		State:        c.GlobalString("from-state"),
		Zip:          c.GlobalString("from-zip"),
		Country:      c.GlobalString("from-country"),
	}
}
