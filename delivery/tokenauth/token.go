package tokenauth

import "errors"

type TokenAuth interface {
	GetToken() string
	VerifyToken(tokenAuth string) error
}

type Token struct {
	Token string
}

func (t *Token) GetToken() string {
	return "123456"
}

func (t *Token) VerifyToken(tokenAuth string) error {
	if tokenAuth != t.GetToken() {
		return errors.New("TOKEN INVALID")
	}
	return nil
}
