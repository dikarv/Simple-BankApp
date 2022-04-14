package usecase

import (
	"enigmacamp.com/bank/repository"
)

type TransferUseCase interface {
	Transfer(SenderAccountNumber, receiverAccountNumber string, amountTransfer int)
}

type transferUseCase struct {
	repo repository.CustomerRepo
}

func (t *transferUseCase) Transfer(SenderAccountNumber, receiverAccountNumber string, amountTransfer int) {
	t.repo.SendTransfer(SenderAccountNumber, amountTransfer)
	t.repo.GetTransfer(receiverAccountNumber, amountTransfer)
}

func NewTransferUseCase(repo repository.CustomerRepo) TransferUseCase {
	return &transferUseCase{
		repo: repo,
	}
}
