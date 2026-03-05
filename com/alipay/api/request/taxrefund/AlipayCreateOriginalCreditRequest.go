package taxrefund

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseTaxRefund "github.com/alipay/global-open-sdk-go/com/alipay/api/response/taxrefund"
)

type AlipayCreateOriginalCreditRequest struct {
	OriginalCreditRequestId string               `json:"originalCreditRequestId,omitempty"`
	PayerAmount             *model.Amount        `json:"payerAmount,omitempty"`
	Payer                   []*model.Merchant    `json:"payer,omitempty"`
	Payee                   *model.TaxRefundUser `json:"payee,omitempty"`
	ScenarioType            string               `json:"scenarioType,omitempty"`
	SubScenarioType         string               `json:"subScenarioType,omitempty"`
	PayerNotificationUrl    string               `json:"payerNotificationUrl,omitempty"`
	Env                     *model.Env           `json:"env,omitempty"`
	TaxRefundFormNumber     string               `json:"taxRefundFormNumber,omitempty"`
	DepartureRegion         string               `json:"departureRegion,omitempty"`
	DeparturePort           string               `json:"departurePort,omitempty"`
	TotalSalesAmount        *model.Amount        `json:"totalSalesAmount,omitempty"`
	Memo                    string               `json:"memo,omitempty"`
}

func NewAlipayCreateOriginalCreditRequest() (*request.AlipayRequest, *AlipayCreateOriginalCreditRequest) {
	alipayCreateOriginalCreditRequest := &AlipayCreateOriginalCreditRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateOriginalCreditRequest, "/aps/api/v1/funds/createOriginalCredit", &responseTaxRefund.AlipayCreateOriginalCreditResponse{})
	return alipayRequest, alipayCreateOriginalCreditRequest
}

func (r *AlipayCreateOriginalCreditRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(r, "/aps/api/v1/funds/createOriginalCredit", &responseTaxRefund.AlipayCreateOriginalCreditResponse{})
}
