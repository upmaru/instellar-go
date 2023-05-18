package instellar

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const slug = "some-test"

const clusterJSON = `
{
  "data": {
    "attributes": {
      "current_state": "%s",
      "id": 8,
      "slug": "%s"
    },
    "id": "8",
    "links": {
      "self": "http://localhost:4000/provision/clusters/8"
    },
    "relationships": {},
    "type": "clusters"
  },
  "included": [],
  "links": {
    "self": "http://localhost:4000/provision/clusters/8"
  }
}
`

func TestGetCluster(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "/provision/clusters/some-test",
		httpmock.NewStringResponder(200, fmt.Sprintf(clusterJSON, "healthy", slug)))

	cluster, _ := client.GetCluster(slug)

	assert.Equal(t, cluster.Data.Attributes.Slug, slug)
}

func TestCreateCluster(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "/provision/clusters",
		httpmock.NewStringResponder(201, fmt.Sprintf(clusterJSON, "connecting", slug)))

	var clusterParams = ClusterParams{
		Name:                           slug,
		Provider:                       "aws",
		CredentialEndpoint:             "https://something:8443",
		CredentialPassword:             "somepass",
		CredentialPasswordConfirmation: "somepass",
	}

	cluster, _ := client.CreateCluster(clusterParams)

	assert.Equal(t, cluster.Data.Attributes.CurrentState, "connecting")
}

func TestUpdateCluster(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PATCH", "/provision/clusters/some-test",
		httpmock.NewStringResponder(200, fmt.Sprintf(clusterJSON, "syncing", slug)))

	cluster, _ := client.UpdateCluster(slug)

	assert.Equal(t, cluster.Data.Attributes.CurrentState, "syncing")
}

func TestDeleteCluster(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "/provision/clusters/some-test",
		httpmock.NewStringResponder(200, fmt.Sprintf(clusterJSON, "deleted", slug)))

	cluster, _ := client.DeleteCluster(slug)

	assert.Equal(t, cluster.Data.Attributes.CurrentState, "deleted")
}
