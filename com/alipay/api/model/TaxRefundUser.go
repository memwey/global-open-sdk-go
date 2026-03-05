package model

type TaxRefundUser struct {
	UserId       string    `json:"userId,omitempty"`
	UserLoginId  string    `json:"userLoginId,omitempty"`
	UserName     *UserName `json:"userName,omitempty"`
}
