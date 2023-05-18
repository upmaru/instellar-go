package instellar

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
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
	Token string `json:"token"`
}

func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
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

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")

	if authToken != nil {
		token := *authToken
		req.Header.Set("Authorization", token)
	}

	reqDump, err := httputil.DumpRequestOut(req, true)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s", string(reqDump))

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d body: %s", res.StatusCode, body)
	}

	return body, err
}
