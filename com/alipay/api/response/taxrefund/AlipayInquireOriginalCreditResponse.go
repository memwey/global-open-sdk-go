package taxrefund

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireOriginalCreditResponse struct {
	response.AlipayResponse
	Result                  *model.Result          `json:"result,omitempty"`
	AcquirerId              string                 `json:"acquirerId,omitempty"`
	PspId                   string                 `json:"pspId,omitempty"`
	OriginalCreditResult    *model.Result          `json:"originalCreditResult,omitempty"`
	ScenarioType            string                 `json:"scenarioType,omitempty"`
	SubScenarioType         string                 `json:"subScenarioType,omitempty"`
	OriginalCreditRequestId string                 `json:"originalCreditRequestId,omitempty"`
	OriginalCreditId        string                 `json:"originalCreditId,omitempty"`
	OriginalCreditTime      string                 `json:"originalCreditTime,omitempty"`
	PayerAmount             *model.Amount          `json:"payerAmount,omitempty"`
	PayeeAmount             *model.Amount          `json:"payeeAmount,omitempty"`
	PayeeQuote              *model.TaxRefundQuote  `json:"payeeQuote,omitempty"`
	Payer                   []*model.Merchant      `json:"payer,omitempty"`
	Payee                   *model.TaxRefundUser   `json:"payee,omitempty"`
}
