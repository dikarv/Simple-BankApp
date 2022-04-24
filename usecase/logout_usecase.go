package usecase

import (
	"enigmacamp.com/bank/repository"
)

type LogoutUseCase interface {
	Logout(accountNumber, token string) error
}

type logoutUseCase struct {
	repo repository.CustomerRepo
}

func (l *logoutUseCase) Logout(accountNumber, token string) error {
	err := l.repo.Logout(accountNumber, token)
	if err != nil {
		return err
	}
	return nil
}

func NewLogoutUseCase(repo repository.CustomerRepo) LogoutUseCase {
	return &logoutUseCase{
		repo: repo,
	}
}
