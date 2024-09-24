package client

import (
	"errors"
	"github.com/OrderMyGear/go-shippo/models"
	"net/http"
)

func (c *Client) CreateSubAccount(input *models.SubAccountInput) (*models.SubAccount, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.SubAccount{}
	err := c.do(http.MethodPost, "/shippo-accounts/", input, output, nil)
	return output, err
}
