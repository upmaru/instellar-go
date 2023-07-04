package instellar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const componentsPath = "provision/components"

type componentReq struct {
	Component ComponentParams `json:"component"`
}

type ComponentParams struct {
	Name       string   `json:"name"`
	Provider   string   `json:"provider"`
	Version    string   `json:"version"`
	Channels   []string `json:"channels"`
	ClusterIDS []int    `json:"cluster_ids"`
	Driver     string   `json:"driver"`
	Credential struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
	}
}

func (c *Client) GetComponent(componentID string) (*Component, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", c.HostURL, componentsPath, componentID), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	component := Component{}

	err = json.Unmarshal(body, &component)

	if err != nil {
		return nil, err
	}

	return &component, nil
}

func (c *Client) CreateComponent(componentParams ComponentParams) (*Component, error) {
	params := componentReq{
		Component: componentParams,
	}

	rb, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s", c.HostURL, componentsPath),
		strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newComponent := Component{}
	err = json.Unmarshal(body, &newComponent)

	if err != nil {
		return nil, err
	}

	return &newComponent, nil
}

func (c *Client) UpdateComponent(componentID string, componentParams ComponentParams) (*Component, error) {
	params := componentReq{
		Component: componentParams,
	}

	rb, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH",
		fmt.Sprintf("%s/%s/%s", c.HostURL, componentsPath, componentID),
		strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	updatedComponent := Component{}
	err = json.Unmarshal(body, &updatedComponent)

	if err != nil {
		return nil, err
	}

	return &updatedComponent, nil
}

func (c *Client) DeleteComponent(componentID string) (*Component, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s", c.HostURL, componentsPath, componentID), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	deletedComponent := Component{}
	err = json.Unmarshal(body, &deletedComponent)

	if err != nil {
		return nil, err
	}

	return &deletedComponent, nil
}
