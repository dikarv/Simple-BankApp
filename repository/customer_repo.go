package repository

import "enigmacamp.com/bank/model"

type CustomerRepo interface {
	Login(customer model.Customer) error
	SendTransfer(accountNumber string, amountTransfer int)
	GetTransfer(accountNumber string, amountTransfer int)
}
