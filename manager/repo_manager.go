package manager

import "enigmacamp.com/goorm/repository"

type RepoManager interface {
	CustomerRepo() repository.CustomerRepository
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.DbConn())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
