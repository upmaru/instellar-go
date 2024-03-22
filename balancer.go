package instellar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const balancersPath = "provision/balancers"
const clusterBalancerPath = "provision/clusters/%s/balancers"

type balancerReq struct {
	Balancer BalancerParams `json:"balancer"`
}

type BalancerParams struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func (c *Client) GetBalancer(balancerID string) (*Balancer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", c.HostURL, balancersPath, balancerID), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	balancer := Balancer{}

	err = json.Unmarshal(body, &balancer)

	if err != nil {
		return nil, err
	}

	return &balancer, nil
}

func (c *Client) CreateBalancer(clusterID string, balancerParams BalancerParams) (*Balancer, error) {
	params := balancerReq{
		Balancer: balancerParams,
	}

	rb, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s", c.HostURL, fmt.Sprintf(clusterBalancerPath, clusterID)),
		strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newBalancer := Balancer{}
	err = json.Unmarshal(body, &newBalancer)

	if err != nil {
		return nil, err
	}

	return &newBalancer, nil
}

func (c *Client) UpdateBalancer(balancerID string, balancerParams BalancerParams) (*Balancer, error) {
	params := balancerReq{
		Balancer: balancerParams,
	}

	rb, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", 
		fmt.Sprintf("%s/%s/%s", c.HostURL, balancersPath, balancerID), 
		strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	updatedBalancer := Balancer{}
	err = json.Unmarshal(body, &updatedBalancer)

	if err != nil {
		return nil, err
	}

	return &updatedBalancer, nil
}

func (c *Client) DeleteBalancer(balancerID string) (*Balancer, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s", c.HostURL, balancersPath, balancerID), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	deletedBalancer := Balancer{}
	err = json.Unmarshal(body, &deletedBalancer)

	if err != nil {
		return nil, err
	}

	return &deletedBalancer, nil
}
