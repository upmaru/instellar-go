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
			"channel_slug": "develop",
			"kit_slug": "%s",
      "installation_id": %d,
			"cluster_id": %d,
			"nodes": ["node-01"]
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
		httpmock.NewStringResponder(200, fmt.Sprintf(uplinkJSON, "active", "pro", 1, 2)))

	uplink, _ := client.GetUplink("8")

	assert.Equal(t, uplink.Data.Attributes.ID, 8)
	assert.Equal(t, uplink.Data.Attributes.Nodes, []string{"node-01"})
}

func TestCreateUplink(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "/provision/clusters/some-cluster/uplinks",
		httpmock.NewStringResponder(200, fmt.Sprintf(uplinkJSON, "created", "pro", 1, 2)))

	var uplinkSetupParams = UplinkSetupParams{
		ChannelSlug: "develop",
		KitSlug:     "pro",
	}

	uplink, _ := client.CreateUplink("some-cluster", uplinkSetupParams)

	assert.Equal(t, uplink.Data.Attributes.CurrentState, "created")
}

func TestUpdateUplink(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PATCH", "/provision/uplinks/8",
		httpmock.NewStringResponder(200, fmt.Sprintf(uplinkJSON, "active", "lite", 1, 2)))

	uplinkSetupParams := UplinkSetupParams{
		ChannelSlug: "develop",
		KitSlug:     "lite",
	}

	uplink, _ := client.UpdateUplink("8", uplinkSetupParams)

	assert.Equal(t, uplink.Data.Attributes.CurrentState, "active")
	assert.Equal(t, uplink.Data.Attributes.KitSlug, "lite")
}

func TestDeleteUplink(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "/provision/uplinks/8",
		httpmock.NewStringResponder(200, fmt.Sprintf(uplinkJSON, "deleted", "pro", 1, 2)))

	uplink, _ := client.DeleteUplink("8")

	assert.Equal(t, uplink.Data.Attributes.CurrentState, "deleted")
}
