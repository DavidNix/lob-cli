package parse

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/davidnix/lob-cli/models"

	cli "gopkg.in/urfave/cli.v1"
)

type headers struct {
	name    int
	address int
	city    int
	state   int
	zip     int
	country int
}

// Addresses returns array of Addresses for csv file
func Addresses(c *cli.Context) ([]models.Address, error) {
	var a = []models.Address{}
	var err error
	f, err := os.Open(c.GlobalString("csv"))
	if err != nil {
		return a, fmt.Errorf("CSV error: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return a, err
	}
	if len(records) < 2 {
		return a, errors.New("no valid csv rows for file " + f.Name())
	}

	h, err := findHeaders(records)
	if err != nil {
		return a, err
	}

	// trim header row
	records = records[1:]

	for _, record := range records {
		a = append(a, models.Address{
			Name:         record[h.name],
			AddressLine1: record[h.address],
			AddressLine2: "",
			City:         record[h.city],
			State:        record[h.state],
			Zip:          record[h.zip],
			Country:      record[h.country],
		})
	}

	return a, nil
}

func findHeaders(records [][]string) (headers, error) {
	row := records[0]

	var indices []int
	for _, v := range []string{
		"name",
		"address",
		"city",
		"state",
		"zip",
		"country",
	} {
		indices = append(indices, indexOf(v, row))
	}

	var h headers
	for _, v := range indices {
		if v < 0 {
			return h, errors.New("invalid header values")
		}
	}

	h = headers{
		name:    indices[0],
		address: indices[1],
		city:    indices[2],
		state:   indices[3],
		zip:     indices[4],
		country: indices[5],
	}

	return h, nil
}

func indexOf(s string, elements []string) int {
	for p, v := range elements {
		if strings.ToLower(v) == s {
			return p
		}
	}
	return -1
}
