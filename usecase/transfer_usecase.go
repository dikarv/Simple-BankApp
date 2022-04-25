package usecase

import (
	"enigmacamp.com/bank/repository"
)

type TransferUseCase interface {
	Transfer(SenderAccountNumber int, receiverAccountNumber int, token string, amountTransfer int, isMerchant bool) error
}

type transferUseCase struct {
	repo repository.CustomerRepo
}

func (t *transferUseCase) Transfer(SenderAccountNumber int, receiverAccountNumber int, token string, amountTransfer int, isMerchant bool) error {
	err := t.repo.SendTransfer(SenderAccountNumber, receiverAccountNumber, token, amountTransfer, isMerchant)
	if err != nil {
		return err
	}
	err = t.repo.GetTransfer(receiverAccountNumber, amountTransfer, isMerchant)
	if err != nil {
		return err
	}
	err = t.repo.AddLogToHistory(SenderAccountNumber, receiverAccountNumber, isMerchant)
	if err != nil {
		return err
	}
	return nil
}

func NewTransferUseCase(repo repository.CustomerRepo) TransferUseCase {
	return &transferUseCase{
		repo: repo,
	}
}
