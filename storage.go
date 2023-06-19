package instellar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const storagesPath = "provision/storages"

type storageReq struct {
	Storage StorageParams `json:"storage"`
}

type StorageParams struct {
	Host                      string `json:"host"`
	Bucket                    string `json:"bucket"`
	Region                    string `json:"region"`
	CredentialAccessKeyID     string `json:"credential_access_key_id"`
	CredentialSecretAccessKey string `json:"credential_secret_access_key"`
}

func (c *Client) GetStorage(storageID string) (*Storage, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", c.HostURL, storagesPath, storageID), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	storage := Storage{}

	err = json.Unmarshal(body, &storage)

	if err != nil {
		return nil, err
	}

	return &storage, nil
}

func (c *Client) CreateStorage(storageParams StorageParams) (*Storage, error) {
	params := storageReq{
		Storage: storageParams,
	}

	rb, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s", c.HostURL, storagesPath),
		strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newStorage := Storage{}
	err = json.Unmarshal(body, &newStorage)

	if err != nil {
		return nil, err
	}

	return &newStorage, nil
}

func (c *Client) UpdateStorage(storageID string, storageParams StorageParams) (*Storage, error) {
	params := storageReq{
		Storage: storageParams,
	}

	rb, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH",
		fmt.Sprintf("%s/%s/%s", c.HostURL, storagesPath, storageID),
		strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	syncingStorage := Storage{}
	err = json.Unmarshal(body, &syncingStorage)

	if err != nil {
		return nil, err
	}

	return &syncingStorage, nil
}

func (c *Client) DeleteStorage(storageID string) (*Storage, error) {
	req, err := http.NewRequest("DELETE",
		fmt.Sprintf("%s/%s/%s", c.HostURL, storagesPath, storageID),
		nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	deletedStorage := Storage{}
	err = json.Unmarshal(body, &deletedStorage)

	if err != nil {
		return nil, err
	}

	return &deletedStorage, nil
}
