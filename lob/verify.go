package lob

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"io/ioutil"

	"github.com/davidnix/lob-cli/models"
)

func (c *Client) VerifyAddress(a models.Address) (models.Address, error) {
	var err error
	body, err := json.Marshal(a)
	if err != nil {
		return a, err
	}

	buf := bytes.NewBuffer(body)
	r, err := http.NewRequest("POST", "https://api.lob.com/v1/us_verifications", buf)
	if err != nil {
		return a, err
	}

	c.config(r)
	resp, err := c.netClient.Do(r)
	// pp, _ := httputil.DumpResponse(resp, true)
	// fmt.Println("RESPONSE:", string(pp))
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return a, fmt.Errorf("verify address %v failed, expected 200, got %v, body: %s", a, resp.Status, body)
	}

	var data struct {
		Address models.Address `json:"address"`
	}
	data.Address = a
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&data); err != nil {
		return a, err
	}

	return data.Address, nil
}
