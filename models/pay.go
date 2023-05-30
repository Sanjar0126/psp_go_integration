package models

type PayRequest struct {
	Environment         string `json:"ENVIRONMENT"`
	VendorID            int64  `json:"VENDOR_ID"`
	PaymentID           int64  `json:"PAYMENT_ID"`
	PaymentName         string `json:"PAYMENT_NAME"`
	AgrTransID          int64  `json:"AGR_TRANS_ID"`
	MerchantTransID     string `json:"MERCHANT_TRANS_ID"`
	MerchantTransAmount int64  `json:"MERCHANT_TRANS_AMOUNT"`
	MerchantTransData   string `json:"MERCHANT_TRANS_DATA"`
	SignTime            int64  `json:"SIGN_TIME"`
	SignString          string `json:"SIGN_STRING"`
}

type PayResponse struct {
	Error         int64  `json:"ERROR"`
	ErrorNote     string `json:"ERROR_NOTE"`
	VendorTransID int64  `json:"VENDOR_TRANS_ID"`
}
