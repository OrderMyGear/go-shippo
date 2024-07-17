package models

import "time"

type LineItem struct {
	Currency           string    `json:"currency,omitempty"`
	ManufactureCountry string    `json:"manufacture_country,omitempty"`
	MaxDeliveryTime    time.Time `json:"max_delivery_time,omitempty"`
	MaxShipTime        time.Time `json:"max_ship_time,omitempty"`
	Quantity           int       `json:"quantity,omitempty"`
	Sku                string    `json:"sku,omitempty"`
	Title              string    `json:"title,omitempty"`
	TotalPrice         string    `json:"total_price,omitempty"`
	VariantTitle       string    `json:"variant_title,omitempty"`
	Weight             string    `json:"weight,omitempty"`
	WeightUnit         string    `json:"weight_unit,omitempty"`
	ObjectID           string    `json:"object_id,omitempty"`
}

// See https://docs.goshippo.com/shippoapi/public-api/#tag/Orders
type OrderInput struct {
	Currency             string       `json:"currency,omitempty"`
	Notes                string       `json:"notes,omitempty"`
	OrderNumber          string       `json:"order_number,omitempty"`
	OrderStatus          string       `json:"order_status,omitempty"`
	PlacedAt             time.Time    `json:"placed_at,omitempty"`
	ShippingCost         string       `json:"shipping_cost,omitempty"`
	ShippingCostCurrency string       `json:"shipping_cost_currency,omitempty"`
	ShippingMethod       string       `json:"shipping_method,omitempty"`
	SubtotalPrice        string       `json:"subtotal_price,omitempty"`
	TotalPrice           string       `json:"total_price,omitempty"`
	TotalTax             string       `json:"total_tax,omitempty"`
	Weight               string       `json:"weight,omitempty"`
	WeightUnit           string       `json:"weight_unit,omitempty"`
	FromAddress          AddressInput `json:"from_address,omitempty"`
	ToAddress            AddressInput `json:"to_address,omitempty"`
	LineItems            []LineItem   `json:"line_items,omitempty"`
}

// See https://goshippo.com/docs/reference#shipments
type Order struct {
	CommonOutputFields
	Currency             string     `json:"currency,omitempty"`
	Notes                string     `json:"notes,omitempty"`
	OrderNumber          string     `json:"order_number,omitempty"`
	OrderStatus          string     `json:"order_status,omitempty"`
	PlacedAt             time.Time  `json:"placed_at,omitempty"`
	ShippingCost         string     `json:"shipping_cost,omitempty"`
	ShippingCostCurrency string     `json:"shipping_cost_currency,omitempty"`
	ShippingMethod       string     `json:"shipping_method,omitempty"`
	SubtotalPrice        string     `json:"subtotal_price,omitempty"`
	TotalPrice           string     `json:"total_price,omitempty"`
	TotalTax             string     `json:"total_tax,omitempty"`
	Weight               string     `json:"weight,omitempty"`
	WeightUnit           string     `json:"weight_unit,omitempty"`
	FromAddress          Address    `json:"from_address,omitempty"`
	ToAddress            Address    `json:"to_address,omitempty"`
	LineItems            []LineItem `json:"line_items,omitempty"`
	ShopApp              string     `json:"shop_app,omitempty"`
	Transactions         []string   `json:"transactions,omitempty"`
}
