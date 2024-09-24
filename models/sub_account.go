package models

type SubAccountInput struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CompanyName string `json:"company_name"`
}

type SubAccount struct {
	SubAccountInput
	CommonOutputFields
	Test       bool        `json:"test"`
	ObjectInfo *ObjectInfo `json:"object_info"`
}
