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
			"endpoint": "127.0.0.1:8443",
			"region": "ap-southeast-1",
			"provider": "aws",
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
		CredentialEndpoint:             "something:8443",
		CredentialPassword:             "somepass",
		CredentialPasswordConfirmation: "somepass",
		InsterraComponentID:            1,
	}

	cluster, _ := client.CreateCluster(clusterParams)

	assert.Equal(t, cluster.Data.Attributes.CurrentState, "connecting")
}

func TestUpdateCluster(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PATCH", "/provision/clusters/some-test",
		httpmock.NewStringResponder(200, fmt.Sprintf(clusterJSON, "syncing", slug)))

	var clusterParams = ClusterParams{
		CredentialEndpoint: "anotherthing:8443",
	}

	cluster, _ := client.UpdateCluster(slug, clusterParams)

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
