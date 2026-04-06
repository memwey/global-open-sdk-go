package taxrefund

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipaySyncTaxRefundFormResponse struct {
	response.AlipayResponse
	Result *model.Result `json:"result,omitempty"`
}

func NewAlipaySyncTaxRefundFormResponse() *AlipaySyncTaxRefundFormResponse {
	return &AlipaySyncTaxRefundFormResponse{}
}
