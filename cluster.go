package instellar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ClusterParams struct {
	Name                           string `json:"name"`
	Provider                       string `json:"provider"`
	CredentialEndpoint             string `json:"credential_endpoint"`
	CredentialPassword             string `json:"credential_password"`
	CredentialPasswordConfirmation string `json:"credential_password_confirmation"`
}

func (c *Client) GetCluster(clusterID string) (*Cluster, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/clusters/%s", c.HostURL, clusterID), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	cluster := Cluster{}

	err = json.Unmarshal(body, &cluster)

	if err != nil {
		return nil, err
	}

	return &cluster, nil
}

func (c *Client) CreateCluster(clusterParams ClusterParams) (*Cluster, error) {
	rb, err := json.Marshal(clusterParams)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/clusters", c.HostURL), strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newCluster := Cluster{}
	err = json.Unmarshal(body, &newCluster)

	if err != nil {
		return nil, err
	}

	return &newCluster, nil
}

func (c *Client) UpdateCluster(clusterID string) (*Cluster, error) {
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/clusters/%s", c.HostURL, clusterID), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	syncingCluster := Cluster{}
	err = json.Unmarshal(body, &syncingCluster)

	if err != nil {
		return nil, err
	}

	return &syncingCluster, nil
}

func (c *Client) DeleteCluster(clusterID string) (*Cluster, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/clusters/%s", c.HostURL, clusterID), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	deletedCluster := Cluster{}
	err = json.Unmarshal(body, &deletedCluster)

	if err != nil {
		return nil, err
	}

	return &deletedCluster, nil
}
