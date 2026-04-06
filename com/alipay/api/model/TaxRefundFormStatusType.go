package model

type TaxRefundFormStatusType string

const (
	TaxRefundFormStatusType_INIT                    TaxRefundFormStatusType = "INIT"
	TaxRefundFormStatusType_STAMPED                 TaxRefundFormStatusType = "STAMPED"
	TaxRefundFormStatusType_REJECTED_BY_CUSTOMS     TaxRefundFormStatusType = "REJECTED_BY_CUSTOMS"
	TaxRefundFormStatusType_RECEIVED                TaxRefundFormStatusType = "RECEIVED"
	TaxRefundFormStatusType_VOIDED                  TaxRefundFormStatusType = "VOIDED"
	TaxRefundFormStatusType_FAILED                  TaxRefundFormStatusType = "FAILED"
	TaxRefundFormStatusType_EXPIRED                 TaxRefundFormStatusType = "EXPIRED"
	TaxRefundFormStatusType_REFUNDED                TaxRefundFormStatusType = "REFUNDED"
	TaxRefundFormStatusType_REFUNDED_NON_ALIPAYPLUS TaxRefundFormStatusType = "REFUNDED_NON_ALIPAYPLUS"
)
