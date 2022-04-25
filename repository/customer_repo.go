package repository

import (
	"errors"
	"strings"
	"time"

	"enigmacamp.com/bank/model"
	"github.com/jmoiron/sqlx"
)

type CustomerRepo interface {
	Login(customer model.Customer) error
	Logout(accountNumber int, token string) error
	SendTransfer(senderAccountNumber int, receiverAccountNumber int, token string, amountTransfer int, isMerchant bool) error
	GetTransfer(accountNumber int, amountTransfer int, isMerchant bool) error
	SaveToken(token string, account_number int) error
	TokenValidator(token string, accountNumber int) error
	AddLogToHistory(senderAccountNumber, receiverAccountNumber int, isMerchant bool) error
	ReceiverExistChecker(accountNumber int, isMerchant bool) error
}

type CustomerRepoImpl struct {
	custDb *sqlx.DB
}

func (c *CustomerRepoImpl) Login(customer model.Customer) error {
	var isUserExists int
	err := c.custDb.Get(&isUserExists, "SELECT COUNT(account_number) FROM customers WHERE account_number = $1 AND user_password = $2", customer.AccountNumber, customer.UserPassword)
	if err != nil {
		return err
	}
	if isUserExists == 0 {
		return errors.New("USER NOT FOUND")
	}
	return nil
}

func (c *CustomerRepoImpl) Logout(accountNumber int, token string) error {
	err := c.TokenValidator(token, accountNumber)
	if err != nil {
		return errors.New("unauthorized userrrs")
	}
	_, err = c.custDb.Exec("UPDATE customers SET token = '' WHERE account_number = $1", accountNumber)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepoImpl) SaveToken(token string, accountNumber int) error {
	_, err := c.custDb.Exec("UPDATE customers SET token = $1 WHERE account_number = $2", token, accountNumber)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepoImpl) SendTransfer(senderAccountNumber int, receiverAccountNumber int, token string, amountTransfer int, isMerchant bool) error {
	err := c.TokenValidator(token, senderAccountNumber)
	if err != nil {
		return errors.New("unauthorized userrrs")
	}
	err = c.ReceiverExistChecker(receiverAccountNumber, isMerchant)
	if err != nil {
		return err
	}
	err = c.BalanceValidator(senderAccountNumber, amountTransfer)
	if err != nil {
		return err
	}
	_, err = c.custDb.Exec("UPDATE customers SET balance = balance - $1 WHERE account_number = $2", amountTransfer, senderAccountNumber)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepoImpl) GetTransfer(accountNumber int, amountTransfer int, isMerchant bool) error {
	if isMerchant {
		_, err := c.custDb.Exec("UPDATE merchants SET balance = balance + $1 WHERE merchantId = $2", amountTransfer, accountNumber)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := c.custDb.Exec("UPDATE customers SET balance = balance + $1 WHERE account_number = $2", amountTransfer, accountNumber)
	if err != nil {
		return err
	}
	return nil
}
func (c *CustomerRepoImpl) ReceiverExistChecker(accountNumber int, isMerchant bool) error {
	var isReceiverExist int
	if isMerchant == true {
		err := c.custDb.Get(&isReceiverExist, "SELECT COUNT(merchantId) FROM merchants WHERE merchantId = $1", accountNumber)
		if err != nil {
			return err
		}
		if isReceiverExist == 0 {
			return errors.New("MERCHANT NOT FOUND")
		}
		return nil
	}
	isReceiverExist = 0
	err := c.custDb.Get(&isReceiverExist, "SELECT COUNT(account_number) FROM customers WHERE account_number = $1", accountNumber)
	if err != nil {
		return err
	}
	if isReceiverExist == 0 {
		return errors.New("ACCOUNT NOT FOUND")
	}
	return nil
}

func (c *CustomerRepoImpl) TokenValidator(token string, accountNumber int) error {
	var selectedToken string
	err := c.custDb.Get(&selectedToken, "SELECT token FROM customers WHERE account_number = $1", accountNumber)
	if err != nil {
		return err
	}
	token = strings.Replace(token, "Bearer ", "", -1)
	if selectedToken != token {
		return errors.New("unauthorized userr")
	}
	return nil
}

func (c *CustomerRepoImpl) BalanceValidator(accounNumber, amountTransfer int) error {
	var balance int
	err := c.custDb.Get(&balance, "SELECT balance FROM customers WHERE account_number = $1", accounNumber)
	if err != nil {
		return err
	}
	if amountTransfer > balance {
		return errors.New("BALANCE IS NOT SUFFICIENT")
	}
	return nil
}

func (c *CustomerRepoImpl) AddLogToHistory(senderAccountNumber, receiverAccountNumber int, isMerchant bool) error {
	time := time.Now().Format("2006-01-02 15:04:05")
	if isMerchant == true {
		_, err := c.custDb.Exec("INSERT INTO history(senderId, receiverMerchantId, successAt) VALUES ($1, $2, $3)", senderAccountNumber, receiverAccountNumber, time)
		if err != nil {
			return errors.New("FAILED LOG")
		}
		return nil
	}
	_, err := c.custDb.Exec("INSERT INTO history(senderId, receiverCustomerId, successAt) VALUES ($1, $2, $3)", senderAccountNumber, receiverAccountNumber, time)
	if err != nil {
		return errors.New("FAILED LOG")
	}
	return nil
}

func NewCustomerRepo(custDb *sqlx.DB) CustomerRepo {
	customerRepo := CustomerRepoImpl{
		custDb: custDb,
	}
	return &customerRepo
}
