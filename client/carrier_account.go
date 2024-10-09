package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/OrderMyGear/go-shippo/models"
)

// CreateCarrierAccount creates a new carrier account object.
func (c *Client) CreateCarrierAccount(input *models.CarrierAccountInput, shippoSubAccountID string) (*models.CarrierAccount, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.CarrierAccount{}
	err := c.do(http.MethodPost, "/carrier_accounts/", input, output, c.subAccountHeader(shippoSubAccountID))
	return output, err
}

func (c *Client) RegisterCarrierAccount(input *models.CarrierAccountInput, shippoSubAccountID string) (*models.CarrierAccount, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.CarrierAccount{}
	err := c.doWithoutVersion(http.MethodPost, "/carrier_accounts/register/new", input, output, c.subAccountHeader(shippoSubAccountID))
	return output, err
}

// RetrieveCarrierAccount retrieves an existing carrier account by object id.
func (c *Client) RetrieveCarrierAccount(objectID string, shippoSubAccountID string) (*models.CarrierAccount, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.CarrierAccount{}
	err := c.do(http.MethodGet, "/carrier_accounts/"+objectID, nil, output, c.subAccountHeader(shippoSubAccountID))
	return output, err
}

// ListAllCarrierAccounts lists all carrier accounts.
func (c *Client) ListAllCarrierAccounts(shippoSubAccountID string) ([]*models.CarrierAccount, error) {
	list := []*models.CarrierAccount{}
	err := c.doList(http.MethodGet, "/carrier_accounts/", nil, func(v json.RawMessage) error {
		item := &models.CarrierAccount{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	}, c.subAccountHeader(shippoSubAccountID))
	return list, err
}

// UpdateCarrierAccount updates an existing carrier account.
// AccountID and Carrier cannot be updated because they form the unique identifier together.
func (c *Client) UpdateCarrierAccount(objectID string, input *models.CarrierAccountInput, shippoSubAccountID string) (*models.CarrierAccount, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.CarrierAccount{}
	err := c.do(http.MethodPut, "/carrier_accounts/"+objectID, input, output, c.subAccountHeader(shippoSubAccountID))
	return output, err
}

func (c *Client) ConnectCarrierAccount(objectID, redirectUrl, state string, shippoSubAccountID string) (string, error) {
	if objectID == "" {
		return "", errors.New("Empty object ID")
	}

	url := fmt.Sprintf("/carrier_accounts/%s/signin/initiate?redirect_uri=%s&state=%s&redirect=false", objectID, redirectUrl, state)

	output := &models.ConnectOauth{}
	err := c.do(http.MethodGet, url, nil, output, c.subAccountHeader(shippoSubAccountID))
	if err != nil {
		return "", err
	}

	return output.RedirectUri, nil
}
