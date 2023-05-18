package instellar

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const jwtToken = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9"

var authJSON string = fmt.Sprintf(`
{
  "data": {
    "token": "%s"
  }
}
`, jwtToken)

var client *Client

func TestNewClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "http://localhost:4000/provision/automation/callback",
		httpmock.NewStringResponder(201, authJSON))

	var host string = "http://localhost:4000"
	var token string = "somekindofstring="

	client, _ = NewClient(&host, &token)

	assert.Equal(t, client.Token, jwtToken, "should be equal")
}
