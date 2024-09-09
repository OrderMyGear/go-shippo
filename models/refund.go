package models

// See https://goshippo.com/docs/reference#refunds
type RefundInput struct {
	Transaction        string `json:"transaction"`
	Async              bool   `json:"async"`
	ShippoSubAccountID string `json:"shippo_sub_account_id,omitempty"`
}

// See https://goshippo.com/docs/reference#refunds
type Refund struct {
	RefundInput
	CommonOutputFields
	Status string `json:"status,omitempty"`
	Test   bool   `json:"test"`
}
