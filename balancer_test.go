package instellar

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const balancerJSON = `
{
	"data": {
		"attributes": {
			"id": 1,
			"name": "some-balancer",
			"address": "some.example.com",
			"current_state": "active"
		},
		"id": "1",
		"links": {
			"self": "http://localhost:4000/provision/balancers/1"
		},
		"relationships": {},
		"type": "balancers"
	},
	"included": [],
	"links": {
		"self": "http://localhost:4000/provision/balancers/1"
	}
}
`

const deletedBalancerJSON = `
{
	"data": {
		"attributes": {
			"id": 1,
			"name": "some-balancer",
			"address": "some.example.com",
			"current_state": "deleted"
		},
		"id": "1",
		"links": {
			"self": "http://localhost:4000/provision/balancers/1"
		},
		"relationships": {},
		"type": "balancers"
	},
	"included": [],
	"links": {
		"self": "http://localhost:4000/provision/balancers/1"
	}
}
`

func TestGetBalancer(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "/provision/balancers/1",
		httpmock.NewStringResponder(200, balancerJSON))

	balancer, _ := client.GetBalancer("1")

	assert.Equal(t, balancer.Data.Attributes.Name, "some-balancer")
}

func TestCreateBalancer(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "/provision/clusters/1/balancers",
		httpmock.NewStringResponder(201, balancerJSON))

	balancerParams := BalancerParams{
		Name:    "some-balancer",
		Address: "some.example.com",
	}

	balancer, _ := client.CreateBalancer("1", balancerParams)

	assert.Equal(t, balancer.Data.Attributes.CurrentState, "active")
}

func TestUpdateBalancer(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PATCH", "/provision/balancers/1",
		httpmock.NewStringResponder(200, balancerJSON))

	balancerParams := BalancerParams{
		Name:    "some-balancer",
		Address: "some.example.com",
	}

	balancer, _ := client.UpdateBalancer("1", balancerParams)

	assert.Equal(t, balancer.Data.Attributes.Name, "some-balancer")
	assert.Equal(t, balancer.Data.Attributes.Address, "some.example.com")
}

func TestDeleteBalancer(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "/provision/balancers/1",
		httpmock.NewStringResponder(200, deletedBalancerJSON))

	balancer, _ := client.DeleteBalancer("1")

	assert.Equal(t, balancer.Data.Attributes.CurrentState, "deleted")
}
