package models

type CancelRequest struct {
	AgrTransID    int64  `json:"AGR_TRANS_ID"`
	VendorTransID int64  `json:"VENDOR_TRANS_ID"`
	SignTime      int64  `json:"SIGN_TIME"`
	SignString    string `json:"SIGN_STRING"`
}
