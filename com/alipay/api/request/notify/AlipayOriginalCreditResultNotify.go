package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

type AlipayOriginalCreditResultNotify struct {
	AlipayNotify
	OriginalCreditResult    *model.Result          `json:"originalCreditResult,omitempty"`
	AcquirerId              string                 `json:"acquirerId,omitempty"`
	PspId                   string                 `json:"pspId,omitempty"`
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
