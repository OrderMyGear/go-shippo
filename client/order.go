package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/OrderMyGear/go-shippo/models"
)

// CreateOrder creates a new order object.
func (c *Client) CreateOrder(input *models.OrderInput) (*models.Order, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.Order{}
	err := c.do(http.MethodPost, "/orders/", input, output)
	return output, err
}

// RetrieveOrder retrieves an existing order by object id.
func (c *Client) RetrieveOrder(objectID string) (*models.Order, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.Order{}
	err := c.do(http.MethodGet, "/orders/"+objectID, nil, output)
	return output, err
}

// ListAllOrders lists all order objects.
func (c *Client) ListAllOrders() ([]*models.Order, error) {
	list := []*models.Order{}
	err := c.doList(http.MethodGet, "/orders/", nil, func(v json.RawMessage) error {
		item := &models.Order{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
