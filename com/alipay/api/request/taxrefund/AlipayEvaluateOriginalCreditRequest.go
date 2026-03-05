package taxrefund

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseTaxRefund "github.com/alipay/global-open-sdk-go/com/alipay/api/response/taxrefund"
)

type AlipayEvaluateOriginalCreditRequest struct {
	PayerAmount      *model.Amount          `json:"payerAmount,omitempty"`
	Payer            []*model.Merchant      `json:"payer,omitempty"`
	PayeeMethod      *model.PaymentMethod   `json:"payeeMethod,omitempty"`
	EvaluationType   string                 `json:"evaluationType,omitempty"`
	ScenarioType     string                 `json:"scenarioType,omitempty"`
	SubScenarioType  string                 `json:"subScenarioType,omitempty"`
	DepartureRegion  string                 `json:"departureRegion,omitempty"`
}

func NewAlipayEvaluateOriginalCreditRequest() (*request.AlipayRequest, *AlipayEvaluateOriginalCreditRequest) {
	alipayEvaluateOriginalCreditRequest := &AlipayEvaluateOriginalCreditRequest{}
	alipayRequest := request.NewAlipayRequest(alipayEvaluateOriginalCreditRequest, "/aps/api/v1/funds/evaluateOriginalCredit", &responseTaxRefund.AlipayEvaluateOriginalCreditResponse{})
	return alipayRequest, alipayEvaluateOriginalCreditRequest
}

func (r *AlipayEvaluateOriginalCreditRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(r, "/aps/api/v1/funds/evaluateOriginalCredit", &responseTaxRefund.AlipayEvaluateOriginalCreditResponse{})
}
