package repository

import "enigmacamp.com/bank/model"

type CustomerRepo interface {
	Login(customer model.Customer) error
	Logout(accountNumber, token string) error
	SendTransfer(accountNumber, token string, amountTransfer int) error
	GetTransfer(accountNumber string, amountTransfer int)
	SaveToken(token, account_number string)
	TokenValidator(token, accountNumber string) error
}
