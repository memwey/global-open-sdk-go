package model

type Passport struct {
	FullName       string `json:"fullName,omitempty"`
	PassportNumber string `json:"passportNumber,omitempty"`
	Nationality    string `json:"nationality,omitempty"`
	ValideDate     string `json:"valideDate,omitempty"`
	ExpireDate     string `json:"expireDate,omitempty"`
	Birthday       string `json:"birthday,omitempty"`
}
