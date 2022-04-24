package appreq

type CustomerRequestLogin struct {
	AccountNumber string `json:"account_number"`
	UserPassword  string `json:"user_password"`
}

type CustomerRequestPayment struct {
	ReceiverAccountNumber string `json:"receiver_account_number"`
	AmountTransfer        int    `json:"amount_transfer"`
	Token                 string `json:"Token"`
}

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}
