package appreq

type CustomerRequestLogin struct {
	AccountNumber string `json:"account_number"`
	UserPassword  string `json:"user_password"`
}

type CustomerRequestPayment struct {
	ReceiverAccountNumber int  `json:"receiver_account_number"`
	AmountTransfer        int  `json:"amount_transfer"`
	IsMerchant            bool `json:"isMerchant"`
}

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}
