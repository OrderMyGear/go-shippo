package shippo

import "github.com/OrderMyGear/go-shippo/client"

const (
	APIVersion20140211 = "2014-02-11"
	APIVersion20161025 = "2016-10-25"
	APIVersion20170329 = "2017-03-29"
	APIVersion20180208 = "2018-02-08"
)

// NewClient creates a new Shippo client.
func NewClient(privateToken string) *client.Client {
	return client.NewClient(privateToken, "")
}

// NewClientWithVersion creates a new Shippo client with API version explicitly specified.
func NewClientWithVersion(privateToken, apiVersion string) *client.Client {
	return client.NewClient(privateToken, apiVersion)
}
