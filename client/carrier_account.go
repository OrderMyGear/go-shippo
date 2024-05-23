package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/OrderMyGear/go-shippo/models"
)

// CreateCarrierAccount creates a new carrier account object.
func (c *Client) CreateCarrierAccount(input *models.CarrierAccountInput) (*models.CarrierAccount, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.CarrierAccount{}
	err := c.do(http.MethodPost, "/carrier_accounts/", input, output)
	return output, err
}

// RetrieveCarrierAccount retrieves an existing carrier account by object id.
func (c *Client) RetrieveCarrierAccount(objectID string) (*models.CarrierAccount, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.CarrierAccount{}
	err := c.do(http.MethodGet, "/carrier_accounts/"+objectID, nil, output)
	return output, err
}

// ListAllCarrierAccounts lists all carrier accounts.
func (c *Client) ListAllCarrierAccounts() ([]*models.CarrierAccount, error) {
	list := []*models.CarrierAccount{}
	err := c.doList(http.MethodGet, "/carrier_accounts/", nil, func(v json.RawMessage) error {
		item := &models.CarrierAccount{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}

// UpdateCarrierAccount updates an existing carrier account.
// AccountID and Carrier cannot be updated because they form the unique identifier together.
func (c *Client) UpdateCarrierAccount(objectID string, input *models.CarrierAccountInput) (*models.CarrierAccount, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.CarrierAccount{}
	err := c.do(http.MethodPut, "/carrier_accounts/"+objectID, input, output)
	return output, err
}

func (c *Client) ConnectCarrierAccount(objectID, redirectUrl, state string) (string, error) {
	if objectID == "" {
		return "", errors.New("Empty object ID")
	}

	url := fmt.Sprintf("/carrier_accounts/%s/signin/initiate?redirect_uri=%s&state=%s", objectID, redirectUrl, state)

	headers := http.Header{}
	err := c.doAndSaveHeaders(http.MethodGet, url, nil, nil, &headers)
	if err != nil {
		return "", err
	}

	c.logger.Println("ConnectCarrierAccount headers", headers)

	return headers.Get("location"), nil
}
