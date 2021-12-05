package parse

import (
	"github.com/DavidNix/lob-cli/models"
	"github.com/urfave/cli/v2"
)

// FromAddress returns parsed return Address
func FromAddress(c *cli.Context) models.Address {
	return models.Address{
		Name:    c.String("from-name"),
		Street:  c.String("from-address"),
		City:    c.String("from-city"),
		State:   c.String("from-state"),
		Zip:     c.String("from-zip"),
		Country: c.String("from-country"),
	}
}
