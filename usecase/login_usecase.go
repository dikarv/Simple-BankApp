package usecase

import (
	"enigmacamp.com/bank/model"
	"enigmacamp.com/bank/repository"
)

type LoginUseCase interface {
	Login(accountNumber int, userPassword, token string) error
}

type loginUseCase struct {
	repo repository.CustomerRepo
}

func (l *loginUseCase) Login(accountNumber int, userPassword, token string) error {
	userAuth := model.NewCustomer(accountNumber, userPassword)
	err := l.repo.Login(userAuth)
	if err != nil {
		return err
	}
	l.repo.SaveToken(token, accountNumber)
	return nil
}

func NewLoginUseCase(repo repository.CustomerRepo) LoginUseCase {
	return &loginUseCase{
		repo: repo,
	}
}
