package instellar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const uplinksPath = "provision/uplinks"
const clusterUplinkPath = "provision/clusters/%s/uplinks"

type uplinkSetupReq struct {
	Uplink UplinkSetupParams `json:"uplink"`
}

type UplinkSetupParams struct {
	Name        string `json:"name"`
	ChannelSlug string `json:"channel_slug"`
	DatabaseURL string `json:"database_url"`
}

func (c *Client) GetUplink(uplinkID string) (*Uplink, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/%s", c.HostURL, uplinksPath, uplinkID),
		nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	uplink := Uplink{}

	err = json.Unmarshal(body, &uplink)

	if err != nil {
		return nil, err
	}

	return &uplink, nil
}

func (c *Client) CreateUplink(clusterID string, uplinkSetupParams UplinkSetupParams) (*Uplink, error) {
	params := uplinkSetupReq{
		Uplink: uplinkSetupParams,
	}

	rb, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s", c.HostURL, fmt.Sprintf(clusterUplinkPath, clusterID)),
		strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newUplinkInstallation := Uplink{}
	err = json.Unmarshal(body, &newUplinkInstallation)

	if err != nil {
		return nil, err
	}

	return &newUplinkInstallation, nil
}

func (c *Client) UpdateUplink(uplinkID string) (*Uplink, error) {
	req, err := http.NewRequest("PATCH",
		fmt.Sprintf("%s/%s/%s", c.HostURL, uplinksPath, uplinkID),
		nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	updatedUplink := Uplink{}
	err = json.Unmarshal(body, &updatedUplink)

	if err != nil {
		return nil, err
	}

	return &updatedUplink, nil
}

func (c *Client) DeleteUplink(uplinkID string) (*Uplink, error) {
	req, err := http.NewRequest("DELETE",
		fmt.Sprintf("%s/%s/%s", c.HostURL, uplinksPath, uplinkID),
		nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	deletedUplink := Uplink{}
	err = json.Unmarshal(body, &deletedUplink)

	if err != nil {
		return nil, err
	}

	return &deletedUplink, nil
}
