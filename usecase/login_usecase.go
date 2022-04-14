package usecase

import (
	"enigmacamp.com/bank/model"
	"enigmacamp.com/bank/repository"
)

type LoginUseCase interface {
	Login(accountNumber string, userPassword string) error
}

type loginUseCase struct {
	repo repository.CustomerRepo
}

func (l *loginUseCase) Login(accountNumber string, userPassword string) error {
	userAuth := model.NewCustomer(accountNumber, userPassword)
	err := l.repo.Login(userAuth)
	if err != nil {
		return err
	}
	return nil
}

func NewLoginUseCase(repo repository.CustomerRepo) LoginUseCase {
	return &loginUseCase{
		repo: repo,
	}
}
