package repository

import (
	"errors"

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

func (c *CustomerRepoImpl) SendTransfer(accountNumber string, amountTransfer int) {
	tx := c.custDb.MustBegin()
	tx.MustExec("UPDATE customers SET balance = balance - $1 WHERE account_number = $2", amountTransfer, accountNumber)
	tx.Commit()
}

func (c *CustomerRepoImpl) GetTransfer(accountNumber string, amountTransfer int) {
	tx := c.custDb.MustBegin()
	tx.MustExec("UPDATE customers SET balance = balance + $1 WHERE account_number = $2", amountTransfer, accountNumber)
	tx.Commit()
}

func NewCustomerRepo(custDb *sqlx.DB) CustomerRepo {
	customerRepo := CustomerRepoImpl{
		custDb: custDb,
	}
	return &customerRepo
}
