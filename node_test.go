package instellar

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const nodeJSON = `
{
  "data": {
    "attributes": {
      "current_state": "%s",
			"slug": "%s",
      "id": 8,
			"public_ip": "127.0.0.1"
    },
    "id": "8",
    "links": {
      "self": "http://localhost:4000/provision/nodes/8"
    },
    "relationships": {},
    "type": "nodes"
  },
  "included": [],
  "links": {
    "self": "http://localhost:4000/provision/nodes/8"
  }
}
`

func TestGetNode(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "/provision/nodes/8",
		httpmock.NewStringResponder(200, fmt.Sprintf(nodeJSON, "healthy", "pizza-node-ham")))

	node, _ := client.GetNode("8")

	assert.Equal(t, node.Data.Attributes.ID, 8)
	assert.Equal(t, node.Data.Attributes.PublicIP, "127.0.0.1")
}

func TestCreateNode(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "/provision/clusters/1/nodes/pizza-node-ham",
		httpmock.NewStringResponder(200, fmt.Sprintf(nodeJSON, "created", "pizza-node-ham")))

	var nodeParams = NodeParams{
		PublicIP: "127.0.0.1",
	}

	node, _ := client.CreateNode("1", "pizza-node-ham", nodeParams)

	assert.Equal(t, node.Data.Attributes.CurrentState, "created")
}

func TestUpdateNode(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "/provision/clusters/1/nodes/pizza-node-ham",
		httpmock.NewStringResponder(200, fmt.Sprintf(nodeJSON, "healthy", "pizza-node-ham")))

	var nodeParams = NodeParams{
		PublicIP: "127.0.0.1",
	}

	node, _ := client.CreateNode("1", "pizza-node-ham", nodeParams)

	assert.Equal(t, node.Data.Attributes.CurrentState, "healthy")
}

func TestDeleteNode(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "/provision/nodes/8",
		httpmock.NewStringResponder(200, fmt.Sprintf(nodeJSON, "deleted", "pizza-node-ham")))

	node, _ := client.DeleteNode("8")

	assert.Equal(t, node.Data.Attributes.CurrentState, "deleted")
}
