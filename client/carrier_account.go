package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/OrderMyGear/go-shippo/models"
)

var (
	ErrEmptyObjectID = errors.New("empty object ID")
	ErrNilInput      = errors.New("nil input")
)

// CreateCarrierAccount creates a new carrier account object.
func (c *Client) CreateCarrierAccount(input *models.CarrierAccountInput, shippoSubAccountID string) (*models.CarrierAccount, error) {
	if input == nil {
		return nil, ErrNilInput
	}

	output := &models.CarrierAccount{}
	err := c.do(http.MethodPost, "/carrier_accounts/", input, output, c.subAccountHeader(shippoSubAccountID))
	return output, err
}

func (c *Client) RegisterCarrierAccount(input *models.CarrierAccountInput, shippoSubAccountID string) (*models.CarrierAccount, error) {
	if input == nil {
		return nil, ErrNilInput
	}

	output := &models.CarrierAccount{}
	err := c.doWithoutVersion(http.MethodPost, "/carrier_accounts/register/new", input, output, c.subAccountHeader(shippoSubAccountID))
	return output, err
}

// RetrieveCarrierAccount retrieves an existing carrier account by object id.
func (c *Client) RetrieveCarrierAccount(objectID string, shippoSubAccountID string) (*models.CarrierAccount, error) {
	if objectID == "" {
		return nil, ErrEmptyObjectID
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
		return nil, ErrEmptyObjectID
	}
	if input == nil {
		return nil, ErrNilInput
	}

	output := &models.CarrierAccount{}
	err := c.do(http.MethodPut, "/carrier_accounts/"+objectID, input, output, c.subAccountHeader(shippoSubAccountID))
	return output, err
}

func (c *Client) ConnectCarrierAccount(objectID, redirectUrl, state string, shippoSubAccountID string) (string, error) {
	if objectID == "" {
		return "", ErrEmptyObjectID
	}

	url := fmt.Sprintf("/carrier_accounts/%s/signin/initiate?redirect_uri=%s&state=%s&redirect=false", objectID, redirectUrl, state)

	output := &models.ConnectOauth{}
	err := c.do(http.MethodGet, url, nil, output, c.subAccountHeader(shippoSubAccountID))
	if err != nil {
		return "", err
	}

	return output.RedirectUri, nil
}

func (c *Client) UploadCarrierAccountDocument(objectID string, input *models.CarrierAccountDocumentInput, shippoSubAccountID string) error {
	if objectID == "" {
		return ErrEmptyObjectID
	}
	if input == nil {
		return ErrNilInput
	}

	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)

	if err := mw.WriteField("document_type", input.DocumentType); err != nil {
		return fmt.Errorf("Error writing document_type field: %s", err.Error())
	}

	fw, err := mw.CreateFormFile("file", input.Filename)
	if err != nil {
		return fmt.Errorf("Error creating file field: %s", err.Error())
	}
	if _, err := io.Copy(fw, input.File); err != nil {
		return fmt.Errorf("Error writing file content: %s", err.Error())
	}

	if err := mw.Close(); err != nil {
		return fmt.Errorf("Error closing multipart writer: %s", err.Error())
	}

	return c.doRaw(http.MethodPost, "/carrier_accounts/"+objectID+"/documents", &buf, mw.FormDataContentType(), nil, c.subAccountHeader(shippoSubAccountID))
}
