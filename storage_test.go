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
			"host": "s3.amazonaws.com",
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
		httpmock.NewStringResponder(200, fmt.Sprintf(storageJSON, "healthy")))

	storage, _ := client.GetStorage("1")

	assert.Equal(t, storage.Data.Attributes.ID, 1)
	assert.Equal(t, storage.Data.Attributes.Host, "s3.amazonaws.com")
}
