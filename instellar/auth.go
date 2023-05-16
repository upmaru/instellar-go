package instellar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) Authenticate() (*AuthResponse, error) {
	if c.Credential.Token == "" {
		return nil, fmt.Errorf("define uid and secret")
	}

	rb, err := json.Marshal(c.Credential)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/provision/automation/callback", c.HostURL), strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
}
