package instellar

import (
	"encoding/json"
	"errors"
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

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/provision/automation/callback", c.HostURL),
		strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, errors.New("unable to login")
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
