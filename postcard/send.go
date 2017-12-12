package postcard

import (
	"fmt"
	"io/ioutil"

	"errors"

	"github.com/davidnix/lob-cli/lob"
	"github.com/davidnix/lob-cli/parse"
	"github.com/fatih/color"
	cli "gopkg.in/urfave/cli.v1"
)

// Send sends postcards from csv of addresses
func Send(c *cli.Context) error {
	var err error
	addresses, err := parse.Addresses(c)
	if err != nil {
		return err
	}

	fmt.Println("Parsed", len(addresses), "addresses")
	if len(addresses) == 0 {
		return errors.New("no addresses found")
	}
	front, back, err := openTemplates(c)
	if err != nil {
		return err
	}

	client := lob.NewClient(c.GlobalString("api-key"))

	fromAddress := parse.FromAddress(c)
	for _, addr := range addresses {
		fmt.Println("\nSending postcard for", addr)
		localErr := client.SendPostcard(fromAddress, addr, front, back)
		if localErr != nil {
			color.Red(fmt.Sprint("Error:", addr, localErr.Error(), "\n"))
		}
	}

	color.Green("\nSending postcards complete!")
	return nil
}

func openTemplates(c *cli.Context) (string, string, error) {
	var err error
	frontURI := c.String("front")
	front, err := ioutil.ReadFile(frontURI)
	if err != nil {
		return "", "", fmt.Errorf("Invalid front file %v: error %v", frontURI, err)
	}

	backURI := c.String("back")
	back, err := ioutil.ReadFile(backURI)
	if err != nil {
		return "", "", fmt.Errorf("Invalid back file %v: error %v", backURI, err)
	}

	return string(front), string(back), nil
}
