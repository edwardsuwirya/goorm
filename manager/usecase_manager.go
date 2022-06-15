package manager

import "enigmacamp.com/goorm/usecase"

type UseCaseManager interface {
	CustomerRegistrationUseCase() usecase.CustomerRegistrationUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) CustomerRegistrationUseCase() usecase.CustomerRegistrationUseCase {
	return usecase.NewCustomerRegistrationUseCase(u.repo.CustomerRepo())
}
func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
