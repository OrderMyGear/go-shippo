package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/OrderMyGear/go-shippo/models"
)

// CreateAddress creates a new address object.
func (c *Client) CreateAddress(input *models.AddressInput, shippoSubAccountID string) (*models.Address, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.Address{}
	err := c.do(http.MethodPost, "/addresses/", input, output, c.subAccountHeader(shippoSubAccountID))
	return output, err
}

// RetrieveAddress retrieves an existing address by object id.
func (c *Client) RetrieveAddress(objectID string, shippoSubAccountID string) (*models.Address, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.Address{}
	err := c.do(http.MethodGet, "/addresses/"+objectID, nil, output, c.subAccountHeader(shippoSubAccountID))
	return output, err
}

// ListAllAddresses lists all addresses.
func (c *Client) ListAllAddresses(shippoSubAccountID string) ([]*models.Address, error) {
	list := []*models.Address{}
	err := c.doList(http.MethodGet, "/addresses/", nil, func(v json.RawMessage) error {
		item := &models.Address{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	}, c.subAccountHeader(shippoSubAccountID))
	return list, err
}
