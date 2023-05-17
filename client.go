package instellar

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
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
	Token string `json:"token"`
}

func NewClient(host, uid *string, secret *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if uid == nil || secret == nil {
		return &c, nil
	}

	data := fmt.Sprintf("organization:%s:%s", *uid, *secret)

	c.Credential = CredentialStruct{
		Token: b64.URLEncoding.EncodeToString([]byte(data)),
	}

	ar, err := c.Authenticate()

	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", token)

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
