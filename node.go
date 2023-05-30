package instellar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const nodesPath = "provision/nodes"
const clusterNodePath = "provision/clusters/%s/nodes/%s"

type nodeReq struct {
	Node NodeParams `json:"node"`
}

type NodeParams struct {
	PublicIP string `json:"public_ip"`
}

func (c *Client) GetNode(nodeID string) (*Node, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", c.HostURL, nodesPath, nodeID), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	node := Node{}

	err = json.Unmarshal(body, &node)

	if err != nil {
		return nil, err
	}

	return &node, nil
}

func (c *Client) CreateNode(clusterID string, nodeSlug string, nodeParams NodeParams) (*Node, error) {
	params := nodeReq{
		Node: nodeParams,
	}

	rb, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT",
		fmt.Sprintf("%s/%s", c.HostURL, fmt.Sprintf(clusterNodePath, clusterID, nodeSlug)),
		strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	upsertedNode := Node{}
	err = json.Unmarshal(body, &upsertedNode)

	if err != nil {
		return nil, err
	}

	return &upsertedNode, nil
}

func (c *Client) UpdateNode(clusterID string, nodeSlug string, nodeParams NodeParams) (*Node, error) {
	params := nodeReq{
		Node: nodeParams,
	}

	rb, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT",
		fmt.Sprintf("%s/%s", c.HostURL, fmt.Sprintf(clusterNodePath, clusterID, nodeSlug)),
		strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	upsertedNode := Node{}
	err = json.Unmarshal(body, &upsertedNode)

	if err != nil {
		return nil, err
	}

	return &upsertedNode, nil
}

func (c *Client) DeleteNode(nodeID string) (*Node, error) {
	req, err := http.NewRequest("DELETE",
		fmt.Sprintf("%s/%s/%s", c.HostURL, nodesPath, nodeID),
		nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	deletedNode := Node{}
	err = json.Unmarshal(body, &deletedNode)

	if err != nil {
		return nil, err
	}

	return &deletedNode, nil
}
