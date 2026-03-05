package taxrefund

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateOriginalCreditResponse struct {
	response.AlipayResponse
	Result             *model.Result          `json:"result,omitempty"`
	AcquirerId         string                 `json:"acquirerId,omitempty"`
	PspId              string                 `json:"pspId,omitempty"`
	OriginalCreditId   string                 `json:"originalCreditId,omitempty"`
	OriginalCreditTime string                 `json:"originalCreditTime,omitempty"`
	PayeeAmount        *model.Amount          `json:"payeeAmount,omitempty"`
	PayeeQuote         *model.TaxRefundQuote  `json:"payeeQuote,omitempty"`
	Payee              *model.TaxRefundUser   `json:"payee,omitempty"`
}
