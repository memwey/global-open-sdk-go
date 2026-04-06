package taxrefund

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseTaxRefund "github.com/alipay/global-open-sdk-go/com/alipay/api/response/taxrefund"
)

type AlipayConfirmOriginalCreditRequest struct {
	OriginalCreditRequestId string `json:"originalCreditRequestId,omitempty"`
	OriginalCreditId        string `json:"originalCreditId,omitempty"`
}

func NewAlipayConfirmOriginalCreditRequest() (*request.AlipayRequest, *AlipayConfirmOriginalCreditRequest) {
	alipayConfirmOriginalCreditRequest := &AlipayConfirmOriginalCreditRequest{}
	alipayRequest := request.NewAlipayRequest(alipayConfirmOriginalCreditRequest, "/aps/api/v1/funds/confirmOriginalCredit", &responseTaxRefund.AlipayConfirmOriginalCreditResponse{})
	return alipayRequest, alipayConfirmOriginalCreditRequest
}

func (r *AlipayConfirmOriginalCreditRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(r, "/aps/api/v1/funds/confirmOriginalCredit", &responseTaxRefund.AlipayConfirmOriginalCreditResponse{})
}
