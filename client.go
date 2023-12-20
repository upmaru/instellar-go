package instellar

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const HostURL string = "https://web.instellar.app"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Credential CredentialStruct
}

type CredentialStruct struct {
	Token string `json:"auth_token"`
}

type AuthResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type AcceptedStates []int

func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if token == nil {
		return &c, nil
	}

	c.Credential = CredentialStruct{
		Token: *token,
	}

	ar, err := c.Authenticate()

	if err != nil {
		return nil, err
	}

	c.Token = ar.Data.Token

	return &c, nil
}

func (arr *AcceptedStates) doesNotContain(target int) bool {
	for _, num := range *arr {
		if num == target {
			return false
		}
	}
	return true
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")

	if c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	accepted := AcceptedStates{http.StatusOK, http.StatusCreated}

	if accepted.doesNotContain(res.StatusCode) {
		return nil, fmt.Errorf("status: %d body: %s", res.StatusCode, body)
	}

	return body, err
}
