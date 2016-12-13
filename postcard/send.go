package postcard

import (
	"fmt"
	"io/ioutil"

	"github.com/davidnix/lob-cli/lob"
	"github.com/davidnix/lob-cli/parse"
	"github.com/fatih/color"
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
	front, back, err := openTemplates(c)
	if err != nil {
		return err
	}

	client := lob.NewClient(c.GlobalString("api-key"))

	fromAddress := parse.FromAddress(c)
	for _, v := range a {
		var localErr error
		verified, localErr := client.VerifyAddress(v)
		if localErr != nil {
			color.Red(localErr.Error())
			continue
		}
		fmt.Println("Sending postcard for", verified)
		localErr = client.SendPostcard(fromAddress, verified, front, back)
		if localErr != nil {
			color.Red(fmt.Sprint("Error:", verified, localErr.Error(), "\n"))
		}
	}

	color.Green("Sending postcards complete!")
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
