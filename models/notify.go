package models

type NotifyRequest struct {
	AgrTransID    int64  `json:"AGR_TRANS_ID"`
	VendorTransID int64  `json:"VENDOR_TRANS_ID"`
	Status        int64  `json:"STATUS"`
	SignTime      int64  `json:"SIGN_TIME"`
	SignString    string `json:"SIGN_STRING"`
}
