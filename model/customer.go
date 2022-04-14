package model

type Customer struct {
	AccountNumber string `db:"account_number"`
	UserPassword  string `db:"user_password"`
	UserBalance   int    `db:"balance"`
}

func (c *Customer) GetAccountNumber() string {
	return c.AccountNumber
}

func (c *Customer) GetUserBalance() int {
	return c.UserBalance
}

func (c *Customer) SetUserPassword(userPassword string) {
	c.UserPassword = userPassword
}

func NewCustomer(accountNumber string, userPassword string) Customer {
	return Customer{
		AccountNumber: accountNumber,
		UserPassword:  userPassword,
	}
}
