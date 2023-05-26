package instellar

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const uplinkJSON = `
{
  "data": {
    "attributes": {
      "current_state": "%s",
      "id": 8,
      "installation_id": %d,
			"cluster_id": %d
    },
    "id": "8",
    "links": {
      "self": "http://localhost:4000/provision/uplinks/8"
    },
    "relationships": {},
    "type": "uplinks"
  },
  "included": [],
  "links": {
    "self": "http://localhost:4000/provision/uplinks/8"
  }
}
`

func TestGetUplink(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "/provision/uplinks/8",
		httpmock.NewStringResponder(200, fmt.Sprintf(uplinkJSON, "active", 1, 2)))

	uplink, _ := client.GetUplink("8")

	fmt.Printf("%+v\n", uplink)

	assert.Equal(t, uplink.Data.Attributes.ID, 8)
}

func TestCreateUplink(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "/provision/clusters/some-cluster/uplinks",
		httpmock.NewStringResponder(200, fmt.Sprintf(uplinkJSON, "created", 1, 2)))

	var uplinkSetupParams = UplinkSetupParams{
		Name:        "test-example",
		ChannelSlug: "develop",
	}

	uplink, _ := client.CreateUplink("some-cluster", uplinkSetupParams)

	assert.Equal(t, uplink.Data.Attributes.CurrentState, "created")
}

func TestUpdateUplink(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PATCH", "/provision/uplinks/8",
		httpmock.NewStringResponder(200, fmt.Sprintf(uplinkJSON, "active", 1, 2)))

	uplink, _ := client.UpdateUplink("8")

	assert.Equal(t, uplink.Data.Attributes.CurrentState, "active")
}

func TestDeleteUplink(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "/provision/uplinks/8",
		httpmock.NewStringResponder(200, fmt.Sprintf(uplinkJSON, "deleted", 1, 2)))

	uplink, _ := client.DeleteUplink("8")

	assert.Equal(t, uplink.Data.Attributes.CurrentState, "deleted")
}
