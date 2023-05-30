package models

type StatusRequest struct {
	AgrTransID int64  `json:"AGR_TRANS_ID"`
	VendorID   int64  `json:"VENDOR_ID"`
	PaymentID  int64  `json:"PAYMENT_ID"`
	SignTime   int64  `json:"SIGN_TIME"`
	SignString string `json:"SIGN_STRING"`
}
