package instellar

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const storageJSON = `
{
	"data": {
		"attributes": {
			"id": 1,
			"current_state": "%s",
			"host": "%s",
			"bucket": "instellar",
			"region": "us-east-1"
		},
		"id": "1",
		"links": {
			"self": "http://localhost:4000/provision/storages/1"
		},
		"relationships": {},
		"type": "storage"
	},
	"included": [],
	"links": {
		"self": "http://localhost:4000/provision/storages/1"
	}
}
`

func TestGetStorage(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "/provision/storages/1",
		httpmock.NewStringResponder(200, fmt.Sprintf(storageJSON, "healthy", "s3.amazonaws.com")))

	storage, _ := client.GetStorage("1")

	assert.Equal(t, storage.Data.Attributes.ID, 1)
	assert.Equal(t, storage.Data.Attributes.Host, "s3.amazonaws.com")
}

func TestCreateStorage(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "/provision/storages",
		httpmock.NewStringResponder(201, fmt.Sprintf(storageJSON, "syncing", "s3.amazonaws.com")))

	var storageParams = StorageParams{
		Host:                      "s3.amazonaws.com",
		Bucket:                    "instellar",
		Region:                    "us-east-1",
		CredentialAccessKeyID:     "something",
		CredentialSecretAccessKey: "secret",
	}

	storage, _ := client.CreateStorage(storageParams)

	assert.Equal(t, storage.Data.Attributes.Host, "s3.amazonaws.com")
	assert.Equal(t, storage.Data.Attributes.Bucket, "instellar")
	assert.Equal(t, storage.Data.Attributes.Region, "us-east-1")
}

func TestUpdateStorage(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PATCH", "/provision/storages/1",
		httpmock.NewStringResponder(200, fmt.Sprintf(storageJSON, "syncing", "object.linode.com")))

	var storageParams = StorageParams{
		Host: "object.linode.com",
	}

	storage, _ := client.UpdateStorage("1", storageParams)

	assert.Equal(t, storage.Data.Attributes.Host, "object.linode.com")
}

func TestDeleteStorage(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "/provision/storages/1",
		httpmock.NewStringResponder(200, fmt.Sprintf(nodeJSON, "deleted", "object.linode.com")))

	storage, _ := client.DeleteStorage("1")

	assert.Equal(t, storage.Data.Attributes.CurrentState, "deleted")
}
