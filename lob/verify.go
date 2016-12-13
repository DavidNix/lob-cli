package lob

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidnix/lob-cli/models"
)

func (c *Client) VerifyAddress(a models.Address) (models.Address, error) {
	var err error
	body, err := json.Marshal(a)
	if err != nil {
		return a, err
	}

	buf := bytes.NewBuffer(body)
	r, err := http.NewRequest("POST", "https://api.lob.com/v1/verify", buf)
	if err != nil {
		return a, err
	}

	c.config(r)
	resp, err := c.netClient.Do(r)
	// pp, _ := httputil.DumpResponse(resp, true)
	// fmt.Println("RESPONSE:", string(pp))
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return a, fmt.Errorf("Verify address %v failed, expected 200, got %v", a, resp.Status)
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
