package taxrefund

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayConfirmOriginalCreditResponse struct {
	response.AlipayResponse
	Result     *model.Result `json:"result,omitempty"`
	AcquirerId string        `json:"acquirerId,omitempty"`
	PspId      string        `json:"pspId,omitempty"`
}
