package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/OrderMyGear/go-shippo/models"
)

// CreateRefund creates a new refund object.
func (c *Client) CreateRefund(input *models.RefundInput, shippoSubAccountID string) (*models.Refund, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.Refund{}
	err := c.do(http.MethodPost, "/refunds/", input, output, c.subAccountHeader(shippoSubAccountID))
	return output, err
}

// RetrieveRefund retrieves an existing refund by object id.
func (c *Client) RetrieveRefund(objectID string, shippoSubAccountID string) (*models.Refund, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.Refund{}
	err := c.do(http.MethodGet, "/refunds/"+objectID, nil, output, c.subAccountHeader(shippoSubAccountID))
	return output, err
}

// ListAllRefunds list all refund objects.
func (c *Client) ListAllRefunds(shippoSubAccountID string) ([]*models.Refund, error) {
	list := []*models.Refund{}
	err := c.doList(http.MethodGet, "/refunds/", nil, func(v json.RawMessage) error {
		item := &models.Refund{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	}, c.subAccountHeader(shippoSubAccountID))
	return list, err
}
