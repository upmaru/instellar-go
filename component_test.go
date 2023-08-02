package instellar

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const componentJSON = `
{
  "data": {
    "attributes": {
      "current_state": "%s",
      "id": 8,
			"provider": "aws",
			"driver": "database/postgresql",
			"version": "15.2",
			"cluster_ids": [1, 2],
			"channels": ["master", "develop"],
			"credential": {
				"username": "postgres",
				"password": "postgres",
				"database": "postgres",
				"host": "localhost",
				"port": 5432
			},
      "slug": "%s"
    },
    "id": "8",
    "links": {
      "self": "http://localhost:4000/provision/components/8"
    },
    "relationships": {},
    "type": "clusters"
  },
  "included": [],
  "links": {
    "self": "http://localhost:4000/provision/components/8"
  }
}
`

func TestGetComponent(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "/provision/components/8",
		httpmock.NewStringResponder(200, fmt.Sprintf(componentJSON, "active", "some-db")))

	component, _ := client.GetComponent("8")

	assert.Equal(t, component.Data.Attributes.Slug, "some-db")
}

func TestCreateComponent(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "/provision/components",
		httpmock.NewStringResponder(201, fmt.Sprintf(componentJSON, "active", "some-db")))

	var componentCredentialParams = ComponentCredentialParams{
		Username: "postgres",
		Password: "postgres",
		Resource: "postgres",
		Host:     "localhost",
		Port:     5432,
	}

	var componentParams = ComponentParams{
		Name:       "some-db",
		Provider:   "aws",
		Driver:     "database/postgresql",
		Version:    "15.2",
		ClusterIDS: []int{1, 2},
		Channels:   []string{"master", "develop"},
		Credential: &componentCredentialParams,
	}

	component, _ := client.CreateComponent(componentParams)

	assert.Equal(t, component.Data.Attributes.Slug, "some-db")
}

func TestUpdateComponent(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PATCH", "/provision/components/8",
		httpmock.NewStringResponder(200, fmt.Sprintf(componentJSON, "active", "some-db")))

	var componentParams = ComponentParams{
		ClusterIDS: []int{2},
	}

	component, _ := client.UpdateComponent("8", componentParams)

	assert.Equal(t, component.Data.Attributes.Slug, "some-db")
}

func TestDeleteComponent(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "/provision/components/8",
		httpmock.NewStringResponder(200, fmt.Sprintf(componentJSON, "deleted", "some-db")))

	component, _ := client.DeleteComponent("8")

	assert.Equal(t, component.Data.Attributes.CurrentState, "deleted")
}
