package manager

import (
	"enigmacamp.com/bank/usecase"
)

type UseCaseManager interface {
	LoginUseCase() usecase.LoginUseCase
	LogoutUseCase() usecase.LogoutUseCase
	TransferUseCase() usecase.TransferUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) LoginUseCase() usecase.LoginUseCase {
	return usecase.NewLoginUseCase(u.repo.CustomerRepo())
}

func (u *useCaseManager) LogoutUseCase() usecase.LogoutUseCase {
	return usecase.NewLogoutUseCase(u.repo.CustomerRepo())
}

func (u *useCaseManager) TransferUseCase() usecase.TransferUseCase {
	return usecase.NewTransferUseCase(u.repo.CustomerRepo())
}

func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
