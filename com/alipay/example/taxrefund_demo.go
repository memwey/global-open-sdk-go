package main

import (
	"fmt"

	defaultAlipayClient "github.com/alipay/global-open-sdk-go/com/alipay/api"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request/taxrefund"
	responseTaxRefund "github.com/alipay/global-open-sdk-go/com/alipay/api/response/taxrefund"
)

func main() {
	const alipayGatewayUrl = ""
	const alipayClientId = ""
	const alipayMerchantPrivateKey = ""
	const alipayAlipayPublicKey = ""

	client := defaultAlipayClient.NewDefaultAlipayClient(
		alipayGatewayUrl,
		alipayClientId,
		alipayMerchantPrivateKey,
		alipayAlipayPublicKey)

	evaluateOriginalCredit(client)
	//createOriginalCredit(client)
	//inquireOriginalCredit(client)
}

func evaluateOriginalCredit(client *defaultAlipayClient.DefaultAlipayClient) {
	alipayRequest, req := taxrefund.NewAlipayEvaluateOriginalCreditRequest()

	req.ScenarioType = "TAX_REFUND"
	req.SubScenarioType = "RESERVATION_TAX_REFUND"

	req.PayerAmount = model.NewAmount("100", "USD")

	req.Payer = []*model.Merchant{
		{
			ReferenceMerchantId: "2188245U41144145",
			MerchantName:        "Merchant Name",
			MerchantMCC:         "5411",
			MerchantAddress: &model.Address{
				Region: "DE",
			},
		},
	}

	req.DepartureRegion = "US"
	req.EvaluationType = "BY_USER_ID"

	req.PayeeMethod = &model.PaymentMethod{
		PaymentMethodType: "CONNECT_WALLET",
		PaymentMethodId:   "11012289272",
	}

	execute, err := client.Execute(alipayRequest)
	if err != nil {
		fmt.Println("SDK error:", err.Error())
		return
	}

	response := execute.(*responseTaxRefund.AlipayEvaluateOriginalCreditResponse)

	switch response.Result.ResultStatus {
	case "S":
		fmt.Println("Success")
		fmt.Println("acquirerId:", response.AcquirerId)
		fmt.Println("pspId:", response.PspId)
		fmt.Println("walletBrandName:", response.WalletBrandName)
		if response.PayeeAmount != nil {
			fmt.Printf("payeeAmount: %s %s\n", response.PayeeAmount.Value, response.PayeeAmount.Currency)
		}
		if response.Payee != nil {
			fmt.Println("payee userId:", response.Payee.UserId)
		}
		if response.Passport != nil {
			fmt.Println("passport holder:", response.Passport.FullName)
		}
	case "F":
		fmt.Println("Failed:", response.Result.ResultCode, response.Result.ResultMessage)
	case "U":
		fmt.Println("Unknown, consider retry:", response.Result.ResultCode, response.Result.ResultMessage)
	}
}

func inquireOriginalCredit(client *defaultAlipayClient.DefaultAlipayClient) {
	alipayRequest, req := taxrefund.NewAlipayInquireOriginalCreditRequest()

	req.OriginalCreditRequestId = "gb_tax_1089760038715669_102775745070000"

	execute, err := client.Execute(alipayRequest)
	if err != nil {
		fmt.Println("SDK error:", err.Error())
		return
	}

	response := execute.(*responseTaxRefund.AlipayInquireOriginalCreditResponse)

	switch response.Result.ResultStatus {
	case "S":
		fmt.Println("Success")
		fmt.Println("originalCreditId:", response.OriginalCreditId)
		fmt.Println("originalCreditTime:", response.OriginalCreditTime)
		if response.OriginalCreditResult != nil {
			fmt.Println("originalCreditResult:", response.OriginalCreditResult.ResultCode, response.OriginalCreditResult.ResultStatus)
		}
		if response.PayeeAmount != nil {
			fmt.Printf("payeeAmount: %s %s\n", response.PayeeAmount.Value, response.PayeeAmount.Currency)
		}
	case "F":
		fmt.Println("Failed:", response.Result.ResultCode, response.Result.ResultMessage)
	case "U":
		fmt.Println("Unknown, consider retry:", response.Result.ResultCode, response.Result.ResultMessage)
	}
}

func createOriginalCredit(client *defaultAlipayClient.DefaultAlipayClient) {
	alipayRequest, req := taxrefund.NewAlipayCreateOriginalCreditRequest()

	req.OriginalCreditRequestId = "gb_tax_1089760038715669_102775745070000"
	req.ScenarioType = "TAX_REFUND"
	req.SubScenarioType = "PORT_INSTANT_TAX_REFUND"

	req.PayerAmount = model.NewAmount("100", "USD")

	req.Payer = []*model.Merchant{
		{
			ReferenceMerchantId: "2188245U41144145",
			MerchantName:        "Merchant Name",
			MerchantMCC:         "5411",
			MerchantAddress: &model.Address{
				Region: "DE",
			},
		},
	}

	req.Payee = &model.TaxRefundUser{
		UserId: "210220900000021958205",
	}

	req.TaxRefundFormNumber = "11048200018287537880"
	req.DepartureRegion = "DE"
	req.DeparturePort = "001"
	req.TotalSalesAmount = model.NewAmount("200", "USD")
	req.PayerNotificationUrl = "https://your-domain.com/notify"

	execute, err := client.Execute(alipayRequest)
	if err != nil {
		fmt.Println("SDK error:", err.Error())
		return
	}

	response := execute.(*responseTaxRefund.AlipayCreateOriginalCreditResponse)

	switch response.Result.ResultStatus {
	case "S":
		fmt.Println("Success")
		fmt.Println("originalCreditId:", response.OriginalCreditId)
		fmt.Println("originalCreditTime:", response.OriginalCreditTime)
		fmt.Println("acquirerId:", response.AcquirerId)
		fmt.Println("pspId:", response.PspId)
		if response.PayeeAmount != nil {
			fmt.Printf("payeeAmount: %s %s\n", response.PayeeAmount.Value, response.PayeeAmount.Currency)
		}
		if response.Payee != nil {
			fmt.Println("payee userId:", response.Payee.UserId)
		}
	case "F":
		fmt.Println("Failed:", response.Result.ResultCode, response.Result.ResultMessage)
	case "U":
		fmt.Println("Unknown, consider retry:", response.Result.ResultCode, response.Result.ResultMessage)
	}
}
