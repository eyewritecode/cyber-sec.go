package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BASEURL = "https://api.shodan.io"

type Client struct {
	apiKey string
}

type ApiInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	HTTPS        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c *Client) ApiStatus() (*ApiInfo, error) {
	endpoint := fmt.Sprintf("%s/api-info?key=%s", BASEURL, c.apiKey)
	res, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var ret ApiInfo

	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
