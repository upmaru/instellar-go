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
	Name                string                     `json:"name,omitempty"`
	Provider            string                     `json:"provider,omitempty"`
	Version             string                     `json:"version,omitempty"`
	Channels            []string                   `json:"channels,omitempty"`
	ClusterIDS          []int                      `json:"cluster_ids,omitempty"`
	Driver              string                     `json:"driver,omitempty"`
	Credential          *ComponentCredentialParams `json:"credential,omitempty"`
	InsterraComponentID int                        `json:"insterra_component_id,omitempty"`
}

type ComponentCredentialParams struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Resource string `json:"resource,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
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
