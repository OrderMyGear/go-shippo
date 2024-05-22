package models

// See https://goshippo.com/docs/reference#carrier-accounts
type CarrierAccountInput struct {
	Carrier    string                 `json:"carrier"`
	AccountID  string                 `json:"account_id"`
	Parameters map[string]interface{} `json:"parameters"`
	Active     bool                   `json:"active"`
}

// See https://goshippo.com/docs/reference#carrier-accounts
type CarrierAccount struct {
	CarrierAccountInput
	CommonOutputFields
	Test       bool        `json:"test"`
	ObjectInfo *ObjectInfo `json:"object_info"`
}

type ObjectInfo struct {
	Authentication *Authentication `json:"authentication"`
}

type Authentication struct {
	Type   string `json:"type"`
	Status string `json:"status"`
}
