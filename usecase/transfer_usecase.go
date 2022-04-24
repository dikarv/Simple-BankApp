package usecase

import (
	"enigmacamp.com/bank/repository"
)

type TransferUseCase interface {
	Transfer(SenderAccountNumber, receiverAccountNumber, token string, amountTransfer int) error
}

type transferUseCase struct {
	repo repository.CustomerRepo
}

func (t *transferUseCase) Transfer(SenderAccountNumber, receiverAccountNumber, token string, amountTransfer int) error {
	err := t.repo.SendTransfer(SenderAccountNumber, token, amountTransfer)
	if err != nil {
		return err
	}
	t.repo.GetTransfer(receiverAccountNumber, amountTransfer)
	return nil
}

func NewTransferUseCase(repo repository.CustomerRepo) TransferUseCase {
	return &transferUseCase{
		repo: repo,
	}
}
