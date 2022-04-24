package repository

import (
	"errors"
	"strings"

	"enigmacamp.com/bank/model"
	"github.com/jmoiron/sqlx"
)

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

func (c *CustomerRepoImpl) Logout(accountNumber, token string) error {
	err := c.TokenValidator(token, accountNumber)
	if err != nil {
		return errors.New("unauthorized userrrs")
	}
	tx := c.custDb.MustBegin()
	tx.MustExec("UPDATE customers SET token = '' WHERE account_number = $1", accountNumber)
	tx.Commit()
	return nil
}

func (c *CustomerRepoImpl) SaveToken(token, accountNumber string) {
	tx := c.custDb.MustBegin()
	tx.MustExec("UPDATE customers SET token = $1 WHERE account_number = $2", token, accountNumber)
	tx.Commit()
}

func (c *CustomerRepoImpl) SendTransfer(accountNumber, token string, amountTransfer int) error {
	err := c.TokenValidator(token, accountNumber)
	if err != nil {
		return errors.New("unauthorized userrrs")
	}
	tx := c.custDb.MustBegin()
	tx.MustExec("UPDATE customers SET balance = balance - $1 WHERE account_number = $2", amountTransfer, accountNumber)
	tx.Commit()
	return nil
}

func (c *CustomerRepoImpl) GetTransfer(accountNumber string, amountTransfer int) {
	tx := c.custDb.MustBegin()
	tx.MustExec("UPDATE customers SET balance = balance + $1 WHERE account_number = $2", amountTransfer, accountNumber)
	tx.Commit()
}

func (c *CustomerRepoImpl) TokenValidator(token, accountNumber string) error {
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

func NewCustomerRepo(custDb *sqlx.DB) CustomerRepo {
	customerRepo := CustomerRepoImpl{
		custDb: custDb,
	}
	return &customerRepo
}
