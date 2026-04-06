package taxrefund

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseTaxRefund "github.com/alipay/global-open-sdk-go/com/alipay/api/response/taxrefund"
)

type AlipaySyncTaxRefundFormRequest struct {
	TaxRefundFormNumber string                        `json:"taxRefundFormNumber,omitempty"`
	FormStatus          model.TaxRefundFormStatusType `json:"formStatus,omitempty"`
	StatusChangeTime    string                        `json:"statusChangeTime,omitempty"`
	FormPrintDate       string                        `json:"formPrintDate,omitempty"`
	FormExpireDate      string                        `json:"formExpireDate,omitempty"`
	TaxRefundAmount     *model.Amount                 `json:"taxRefundAmount,omitempty"`
	Merchants           []*model.Merchant             `json:"merchants,omitempty"`
	UserId              string                        `json:"userId,omitempty"`
	Memo                string                        `json:"memo,omitempty"`
}

func NewAlipaySyncTaxRefundFormRequest() (*request.AlipayRequest, *AlipaySyncTaxRefundFormRequest) {
	alipaySyncTaxRefundFormRequest := &AlipaySyncTaxRefundFormRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySyncTaxRefundFormRequest, "/aps/api/v1/funds/syncTaxRefundForm", &responseTaxRefund.AlipaySyncTaxRefundFormResponse{})
	return alipayRequest, alipaySyncTaxRefundFormRequest
}

func (r *AlipaySyncTaxRefundFormRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(r, "/aps/api/v1/funds/syncTaxRefundForm", &responseTaxRefund.AlipaySyncTaxRefundFormResponse{})
}
