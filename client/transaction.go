package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/OrderMyGear/go-shippo/models"
)

// PurchaseShippingLabel creates a new transaction object and purchases the shipping label for the provided rate.
func (c *Client) PurchaseShippingLabel(input *models.TransactionInput) (*models.Transaction, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.Transaction{}
	err := c.do(http.MethodPost, "/transactions/", input, output)
	return output, err
}

// RetrieveTransaction retrieves an existing transaction by object id.
func (c *Client) RetrieveTransaction(objectID string, shippoSubAccountID string) (*models.Transaction, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.Transaction{}
	err := c.do(http.MethodGet, "/transactions/"+objectID, &models.ShippoSubAccount{ShippoSubAccountID: shippoSubAccountID}, output)
	return output, err
}

// ListAllTransactions lists all transaction objects.
func (c *Client) ListAllTransactions() ([]*models.Transaction, error) {
	list := []*models.Transaction{}
	err := c.doList(http.MethodGet, "/transactions/", nil, func(v json.RawMessage) error {
		item := &models.Transaction{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
