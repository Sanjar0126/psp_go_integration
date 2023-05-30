package models

type StatementRequest struct {
	From       int64  `json:"FROM"`
	To         int64  `json:"TO"`
	SignTime   int64  `json:"SIGN_TIME"`
	SignString string `json:"SIGN_STRING"`
}

type StatementTransaction struct {
	Environment         string `json:"ENVIRONMENT"`
	AgrTransID          string `json:"AGR_TRANS_ID"`
	VendorTransID       string `json:"VENDOR_TRANS_ID"`
	MerchantTransID     string `json:"MERCHANT_TRANS_ID"`
	MerchantTransAmount string `json:"MERCHANT_TRANS_AMOUNT"`
	State               string `json:"STATE"`
	Date                string `json:"DATE"`
}

type StatementResponse struct {
	Error        int64                  `json:"ERROR"`
	ErrorNote    string                 `json:"ERROR_NOTE"`
	Transactions []StatementTransaction `json:"TRANSACTIONS"`
}
