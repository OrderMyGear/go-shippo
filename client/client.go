package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/OrderMyGear/go-shippo/errors"
	"github.com/OrderMyGear/go-shippo/models"
)

const (
	ShippoAccountIdHeader     = "SHIPPO-ACCOUNT-ID"
	shippoAPIBaseURL          = "https://api.goshippo.com/v1"
	shippoAPIBaseURLNoVersion = "https://api.goshippo.com"
)

type Client struct {
	privateToken string
	apiVersion   string
	logger       *log.Logger
}

type listOutputCallback func(v json.RawMessage) error

// NewClient creates a new Shippo API client instance.
func NewClient(privateToken, apiVersion string) *Client {
	return &Client{
		privateToken: privateToken,
		apiVersion:   apiVersion,
	}
}

// SetTraceLogger sets a new trace logger and returns the old logger.
// If logger is not nil, Client will output all internal messages to the logger.
func (c *Client) SetTraceLogger(logger *log.Logger) *log.Logger {
	oldLogger := c.logger
	c.logger = logger
	return oldLogger
}

func (c *Client) do(method, path string, input, output interface{}, headers map[string]string) error {
	url := getBaseUrl(headers) + path

	req, err := c.createRequest(method, url, input, headers)
	if err != nil {
		return fmt.Errorf("Error creating request object: %s", err.Error())
	}

	if err := c.executeRequest(req, output); err != nil {
		if aerr, ok := err.(*errors.APIError); ok {
			return aerr
		}
		return fmt.Errorf("Error executing request: %s", err.Error())
	}

	return nil
}

func getBaseUrl(headers map[string]string) string {
	if headers != nil {
		if _, ok := headers[ShippoAccountIdHeader]; ok {
			return shippoAPIBaseURLNoVersion
		}
	}
	return shippoAPIBaseURL
}

func (c *Client) doList(method, path string, input interface{}, outputCallback listOutputCallback, headers map[string]string) error {
	nextURL := getBaseUrl(headers) + path + "?results=25"

	for {
		req, err := c.createRequest(method, nextURL, input, headers)
		if err != nil {
			return fmt.Errorf("Error creating request object: %s", err.Error())
		}

		listOutput := &models.ListAPIOutput{}
		if err := c.executeRequest(req, listOutput); err != nil {
			if aerr, ok := err.(*errors.APIError); ok {
				return aerr
			}
			return fmt.Errorf("Error executing request: %s", err.Error())
		}

		for _, v := range listOutput.Results {
			if err := outputCallback(v); err != nil {
				return fmt.Errorf("Error unmarshalling output item: %s", err.Error())
			}
		}

		if listOutput.NextPageURL == nil {
			break
		}

		nextURL = *listOutput.NextPageURL
	}

	return nil
}

func (c *Client) createRequest(method, url string, bodyObject interface{}, headers map[string]string) (req *http.Request, err error) {
	var reqBodyDebug []byte

	if c.logger != nil {
		defer func() {
			if err != nil {
				c.logPrintf("Client.createRequest() error: %s", err.Error())
				return
			} else if req == nil {
				c.logPrintf("Client.createRequest() req=nil, error=nil")
				return
			}

			for hk, hva := range req.Header {
				if hk != "Authorization" {
					for _, hv := range hva {
						c.logPrintf("Client.createRequest() Header %s=%s", hk, hv)
					}
				}
			}

			body := ""
			if reqBodyDebug != nil {
				body = string(reqBodyDebug)
			}

			c.logPrintf("Client.createRequest() HTTP request created: method=%q, url=%q, body=%q",
				req.Method, req.URL.String(), body)
		}()
	}

	var reqBody io.Reader
	if bodyObject != nil {
		data, err := json.Marshal(bodyObject)
		if err != nil {
			return nil, fmt.Errorf("Error marshaling body object: %s", err.Error())
		}

		reqBodyDebug = data

		reqBody = bytes.NewBuffer(data)
	}

	req, err = http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("Error creating HTTP request: %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "ShippoToken "+c.privateToken)
	if c.apiVersion != "" {
		req.Header.Set("Shippo-API-Version", c.apiVersion)
	}

	// no keep-alive
	req.Header.Set("Connection", "close")
	req.Close = true

	// add any passed in headers
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}

func (c *Client) executeRequest(req *http.Request, output interface{}) (err error) {
	if c.logger != nil {
		defer func() {
			if err != nil {
				c.logPrintf("Client.executeRequest() error: %s", err.Error())
			}
		}()
	}

	httpClient := http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error making HTTP request: %s", err.Error())
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body data: %s", err.Error())
	}

	if c.logger != nil {
		c.logPrintf("Client.executeRequest() response: status=%q, body=%q", res.Status, string(resData))
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		if output != nil && len(resData) > 0 {
			if err := json.Unmarshal(resData, output); err != nil {
				return fmt.Errorf("Error unmarshaling response data: %s", err.Error())
			}
		}

		return nil
	}

	return &errors.APIError{
		Status:       res.StatusCode,
		ResponseBody: resData,
	}
}

func (c *Client) logPrintf(format string, args ...interface{}) {
	if c.logger != nil {
		c.logger.Printf(format, args...)
	}
}

func (c *Client) subAccountHeader(id string) map[string]string {
	headers := make(map[string]string)
	if len(id) > 0 {
		headers[ShippoAccountIdHeader] = id
	}
	return headers
}
