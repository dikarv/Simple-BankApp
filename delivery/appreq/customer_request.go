package appreq

type CustomerRequestLogin struct {
	AccountNumber string `json:"account_number"`
	UserPassword  string `json:"user_password"`
}

type CustomerRequestPayment struct {
	SenderAccountNumber   string `json:"sender_account_number"`
	ReceiverAccountNumber string `json:"receiver_account_number"`
	AmountTransfer        int    `json:"amount_transfer"`
	Token                 string `json:"Token"`
}
