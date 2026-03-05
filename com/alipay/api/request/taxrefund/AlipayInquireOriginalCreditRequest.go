package taxrefund

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseTaxRefund "github.com/alipay/global-open-sdk-go/com/alipay/api/response/taxrefund"
)

type AlipayInquireOriginalCreditRequest struct {
	OriginalCreditRequestId string `json:"originalCreditRequestId,omitempty"`
	OriginalCreditId        string `json:"originalCreditId,omitempty"`
}

func NewAlipayInquireOriginalCreditRequest() (*request.AlipayRequest, *AlipayInquireOriginalCreditRequest) {
	alipayInquireOriginalCreditRequest := &AlipayInquireOriginalCreditRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireOriginalCreditRequest, "/aps/api/v1/funds/inquireOriginalCredit", &responseTaxRefund.AlipayInquireOriginalCreditResponse{})
	return alipayRequest, alipayInquireOriginalCreditRequest
}

func (r *AlipayInquireOriginalCreditRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(r, "/aps/api/v1/funds/inquireOriginalCredit", &responseTaxRefund.AlipayInquireOriginalCreditResponse{})
}
