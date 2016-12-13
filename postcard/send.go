package postcard

import (
	"fmt"

	"github.com/davidnix/lob-cli/lob"
	"github.com/davidnix/lob-cli/parse"
	cli "gopkg.in/urfave/cli.v1"
)

// Send sends postcards from csv of addresses
func Send(c *cli.Context) error {
	var err error
	a, err := parse.Addresses(c)
	if err != nil {
		return err
	}

	fmt.Println("Parsed", len(a), "addresses")
	if len(a) == 0 {
		return nil
	}
	client := lob.NewClient(c.GlobalString("api-key"))

	var errors []error

	for _, v := range a {
		verified, verifyErr := client.VerifyAddress(v)
		if verifyErr != nil {
			errors = append(errors, err)
		} else {
			// process verified
		}
	}

	return nil
}
