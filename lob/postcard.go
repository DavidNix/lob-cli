package lob

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/davidnix/lob-cli/models"
)

func (c *Client) SendPostcard(from, to models.Address, front, back string) error {
	data := url.Values{}
	data.Set("size", "4x6")

	data.Set("front", front)
	data.Set("back", back)

	data.Set("to[name]", to.Name)
	data.Set("to[address_line1]", to.AddressLine1)
	data.Set("to[address_line2]", to.AddressLine2)
	data.Set("to[address_city]", to.City)
	data.Set("to[address_state]", to.State)
	data.Set("to[address_zip]", to.Zip)
	data.Set("to[address_country]", to.Country)

	data.Set("from[name]", from.Name)
	data.Set("from[address_line1]", from.AddressLine1)
	data.Set("from[address_line2]", from.AddressLine2)
	data.Set("from[address_city]", from.City)
	data.Set("from[address_state]", from.State)
	data.Set("from[address_zip]", from.Zip)
	data.Set("from[address_country]", from.Country)

	var err error
	body := bytes.NewBufferString(data.Encode())
	r, err := http.NewRequest("POST", "https://api.lob.com/v1/postcards", body)
	if err != nil {
		return err
	}

	c.config(r)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := c.netClient.Do(r)
	if err != nil {
		return fmt.Errorf("Send postcard for %v failed error:", to, err)
	}
	// pp, _ := httputil.DumpResponse(resp, true)
	// fmt.Println("RESPONSE:", string(pp))
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Send postcard for %v failed, expected 200, got %v", to, resp.Status)
	}

	var apiResponse struct {
		Delivery string `json:"expected_delivery_date"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&apiResponse); err != nil {
		fmt.Println("Unable to determine delivery date")
	} else {
		fmt.Println("Expected delivery date", apiResponse.Delivery)
	}

	return nil
}