package manager

import "enigmacamp.com/bank/repository"

type RepoManager interface {
	CustomerRepo() repository.CustomerRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
