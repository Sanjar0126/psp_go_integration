package models

type GetInfoRequest struct {
	MerchantTransID string `json:"MERCHANT_TRANS_ID"`
	SignTime        int64  `json:"SIGN_TIME"`
	SignString      string `json:"SIGN_STRING"`
}

type GetInfoParams struct {
	FullName string `json:"full_name"`
	Balance  string `json:"balance"`
	Email    string `json:"email"`
}

type GetInfoResponse struct {
	Error      int64         `json:"ERROR"`
	ErrorNote  string        `json:"ERROR_NOTE"`
	Parameters GetInfoParams `json:"PARAMETERS"`
}
