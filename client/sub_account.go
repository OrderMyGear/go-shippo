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

	// Truncate first and last name to 30 characters to avoid Stripe API errors
	if len(input.FirstName) > 30 {
		input.FirstName = input.FirstName[:30]
	}
	if len(input.LastName) > 30 {
		input.LastName = input.LastName[:30]
	}

	output := &models.SubAccount{}
	err := c.doWithoutVersion(http.MethodPost, "/shippo-accounts", input, output, nil)
	return output, err
}
